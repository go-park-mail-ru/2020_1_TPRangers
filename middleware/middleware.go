package middleware

import (
	"../models"
	"github.com/labstack/echo"
	"net/http"
)

func SetCorsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		//TODO: убрать из корса
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")

		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, POST")
		c.Response().Header().Set("Set-Cookie", "HttpOnly, Secure, SameSite=Strict")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, X-Login, Set-Cookie, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, csrf-token, Authorization")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Vary", "Cookie")

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusOK)
		}

		return next(c)

	}
}

func PanicMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		defer func() error {
			if err := recover(); err != nil {
				return c.JSON(http.StatusInternalServerError, models.JsonStruct{Err: "server panic ! "})
			}
			return next(c)
		}()
		return next(c)

	}
}
