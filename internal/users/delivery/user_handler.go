package delivery

import (
	"../../tools/errors"
	"../../models"
	"../../users"
	"../../users/usecase"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserDeliveryRealisation struct {
	userLogic users.UserUseCase
	logger    *zap.SugaredLogger
}

func (userD UserDeliveryRealisation) GetUser(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")
	login := rwContext.Param("id")
	userData, err := userD.userLogic.GetUser(login)

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
		)

		return rwContext.JSON(http.StatusNotFound, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: userData})
}

func (userD UserDeliveryRealisation) Profile(rwContext echo.Context) error {
	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", err.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	userProfile, err := userD.userLogic.Profile(cookie.Value)

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: userProfile})
}

func (userD UserDeliveryRealisation) GetSettings(rwContext echo.Context) error {
	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	userSettings, err := userD.userLogic.GetSettings(cookie.Value)

	respErrStat := 0

	switch err {
	case errors.InvalidCookie:

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		respErrStat = http.StatusUnauthorized
	case errors.FailReadFromDB:
		respErrStat = http.StatusInternalServerError
	}

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", respErrStat),
		)

		return rwContext.JSON(respErrStat, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: userSettings})
}

func (userD UserDeliveryRealisation) UploadSettings(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newUserSettings := new(models.Settings)

	err = rwContext.Bind(newUserSettings)

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: errors.FailDecode.Error()})
	}

	userSettings, err := userD.userLogic.UploadSettings(cookie.Value, *newUserSettings)

	respErrStat := 0

	switch err {
	case errors.InvalidCookie:
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		respErrStat = http.StatusUnauthorized
	case errors.FailReadFromDB:
		respErrStat = http.StatusInternalServerError
	}

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", respErrStat),
		)

		return rwContext.JSON(respErrStat, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: userSettings})

}

func (userD UserDeliveryRealisation) Login(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	userAuthData := new(models.Auth)

	err := rwContext.Bind(userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	info, _ := uuid.NewV4()
	exprTime := 12 * time.Hour
	cookieValue := info.String()

	err = userD.userLogic.Login(*userAuthData, cookieValue, exprTime)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) Logout(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return rwContext.NoContent(http.StatusOK)
	}

	err = userD.userLogic.Logout(cookie.Value)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	rwContext.SetCookie(cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) Register(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	userAuthData := new(models.Register)

	err := rwContext.Bind(userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	info, _ := uuid.NewV4()
	exprTime := 12 * time.Hour
	cookieValue := info.String()

	err = userD.userLogic.Register(*userAuthData, cookieValue, exprTime)

	errResStatus := 0

	switch err {
	case errors.AlreadyExist:
		errResStatus = http.StatusConflict
	case errors.FailReadFromDB:
		errResStatus = http.StatusInternalServerError
	}

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errResStatus),
		)

		return rwContext.NoContent(errResStatus)
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	cookie := http.Cookie{
		Name:    "session_id",
		Value:   cookieValue,
		Expires: time.Now().Add(exprTime),
	}
	rwContext.SetCookie(&cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) AddFriend(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	friendLogin := rwContext.Param("id")
	err = userD.userLogic.AddFriend(cookie.Value, friendLogin)

	errRespStatus := 0

	switch err {
	case errors.InvalidCookie:

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		errRespStatus = http.StatusUnauthorized

	case errors.FailAddFriend:
		errRespStatus = http.StatusInternalServerError
	}

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStatus),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewUserDelivery(log *zap.SugaredLogger, userRealisation usecase.UserUseCaseRealisation) UserDeliveryRealisation {
	return UserDeliveryRealisation{userLogic: userRealisation, logger: log}
}

func (userD UserDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/login", userD.Login)
	server.POST("/api/v1/registration", userD.Register)

	server.PUT("/api/v1/settings", userD.UploadSettings)
	server.PUT("/api/v1/user/:id", userD.AddFriend)

	server.GET("/api/v1/profile", userD.Profile)
	server.GET("/api/v1/settings", userD.GetSettings)
	server.GET("/api/v1/user/:id", userD.GetUser)

	server.DELETE("/api/v1/auth", userD.Logout)
}