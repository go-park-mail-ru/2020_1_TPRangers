package middleware

import (
	"context"
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"main/internal/csrf"
	sessions "main/internal/microservices/authorization/delivery"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
	"time"
)

type MiddlewareHandler struct {
	logger      *zap.SugaredLogger
	sessChecker sessions.SessionCheckerClient
	httpOrigin  string
	tracker     *prometheus.CounterVec
}

func NewMiddlewareHandler(logger *zap.SugaredLogger, checker sessions.SessionCheckerClient, metric *prometheus.CounterVec, origin string) MiddlewareHandler {
	return MiddlewareHandler{logger: logger, sessChecker: checker, httpOrigin: origin, tracker: metric}
}

func (mh MiddlewareHandler) SetMiddleware(server *echo.Echo) {
	server.Use(mh.SetCorsMiddleware)

	logFunc := mh.AccessLog()
	server.Use(mh.PanicMiddleWare)
	authFunc := mh.CheckAuthentication()
	csrfFunc := mh.CSRF()

	server.Use(authFunc)
	server.Use(logFunc)

	server.Use(csrfFunc)
}

func (mh MiddlewareHandler) SetCorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Response().Header().Set("Access-Control-Allow-Origin", mh.httpOrigin)
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-Csrf-Token, csrf-token, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Vary", "Cookie")

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}

		return next(c)

	}
}

func (mh MiddlewareHandler) PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		defer func() error {
			if err := recover(); err != nil {
				rId := c.Get("REQUEST_ID").(string)
				mh.logger.Info(
					zap.String("ID", rId),
					zap.String("ERROR" , err.(error).Error()),
					zap.Int("ANSWER STATUS", http.StatusInternalServerError),
				)
				mh.tracker.WithLabelValues(strconv.Itoa(http.StatusInternalServerError), c.Request().URL.Path, c.Request().Method, "0").Inc()
				return c.JSON(http.StatusInternalServerError, models.JsonStruct{Err: "server panic ! "})
			}
			return nil
		}()

		return next(c)
	}
}

func (mh MiddlewareHandler) AccessLog() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(rwContext echo.Context) error {

			uniqueID := uuid.NewV4()
			start := time.Now()
			rwContext.Set("REQUEST_ID", uniqueID.String())

			mh.logger.Info(
				zap.String("ID", uniqueID.String()),
				zap.String("URL", rwContext.Request().URL.Path),
				zap.String("METHOD", rwContext.Request().Method),
			)

			err := next(rwContext)

			respTime := time.Since(start)
			mh.logger.Info(
				zap.String("ID", uniqueID.String()),
				zap.Duration("TIME FOR ANSWER", respTime),
			)

			if rwContext.Request().URL.Path != "/metrics" {
				mh.tracker.WithLabelValues(strconv.Itoa(rwContext.Response().Status), rwContext.Request().URL.Path, rwContext.Request().Method, respTime.String()).Inc()
			}

			return err

		}
	}
}

func (mh MiddlewareHandler) CheckAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(rwContext echo.Context) error {

			cookie, err := rwContext.Cookie("session_id")

			userId := &sessions.UserId{
				UserId: -1,
			}

			if err == nil {
				userId, err = mh.sessChecker.CheckSession(context.Background(), &sessions.SessionData{
					Cookies: cookie.Value,
				})
			}

			if err != nil {
				cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
				rwContext.SetCookie(cookie)
			}

			rwContext.Set("user_id", int(userId.UserId))
			return next(rwContext)

		}
	}
}

func (mh MiddlewareHandler) CSRF() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(rwContext echo.Context) error {
			if rwContext.Request().RequestURI == "/api/v1/settings" || rwContext.Request().Method == "PUT" {
				cookie, err := rwContext.Cookie("session_id")
				if err != nil {
					mh.logger.Debug(
						zap.String("COOKIE", errors.CookieExpired.Error()),
						zap.Int("ANSWER STATUS", http.StatusUnauthorized),
					)

					return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
				}

				tokenReq := rwContext.Request().Header.Get("X-CSRF-Token")

				isValidCsrf, err := csrf.Tokens.Check(cookie.Value, tokenReq)

				if err != nil {
					return rwContext.JSON(http.StatusForbidden, models.JsonStruct{Err: errors.CookieExpired.Error()})
				}

				if isValidCsrf == false {
					return rwContext.JSON(http.StatusForbidden, models.JsonStruct{Err: errors.CookieExpired.Error()})
				}
			}
			return next(rwContext)
		}
	}
}
