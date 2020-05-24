package delivery

import (
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"main/internal/csrf"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/internal/users"
	"net/http"
	"time"
)

type UserDeliveryRealisation struct {
	userLogic users.UserUseCase
	logger    *zap.SugaredLogger
}

func (userD UserDeliveryRealisation) GetUser(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)
	login := rwContext.Param("id")

	userId := rwContext.Get("user_id").(int)

	fmt.Println("GET USE DATA : ", uId, login, userId)

	var userData models.OtherUserProfileData
	var err error

	if userId != -1 {
		userData, err = userD.userLogic.GetUserProfileWhileLogged(login, userId)
		if err != nil {
			return err
		}
		userData.IsFriends, err = userD.userLogic.CheckFriendship(userId, login)
	} else {
		userData, err = userD.userLogic.GetOtherUserProfileNotLogged(login)
	}

	fmt.Println(userData)

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

	return rwContext.JSON(http.StatusOK, userData)
}

func (userD UserDeliveryRealisation) Profile(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	userProfile, err := userD.userLogic.GetMainUserProfile(userId)

	if err != nil {

		userD.logger.Info(
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

	return rwContext.JSON(http.StatusOK, userProfile)
}
func (userD UserDeliveryRealisation) GetSettings(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	userSettings, err := userD.userLogic.GetSettings(userId)

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, userSettings)
}

func (userD UserDeliveryRealisation) UploadSettings(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	b, err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	newUserSettings := new(models.Settings)

	err = newUserSettings.UnmarshalJSON(b)

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: errors.FailDecode.Error()})
	}

	userSettings, err := userD.userLogic.UploadSettings(userId, *newUserSettings)

	respErrStat := 0

	switch err {
	case errors.FailReadFromDB:
		respErrStat = http.StatusInternalServerError
	case errors.FailSendToDB:
		respErrStat = http.StatusConflict
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

	return rwContext.JSON(http.StatusOK, userSettings)

}

func (userD UserDeliveryRealisation) Login(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userAuthData := new(models.Auth)

	b, err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = userAuthData.UnmarshalJSON(b)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	exprTime := 12 * time.Hour

	cookieValue, err := userD.userLogic.Login(*userAuthData)

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

	uId := rwContext.Get("REQUEST_ID").(string)

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
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	rwContext.SetCookie(cookie)

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) Register(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userAuthData := new(models.Register)

	b, err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = userAuthData.UnmarshalJSON(b)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	exprTime := 12 * time.Hour

	cookieValue, err := userD.userLogic.Register(*userAuthData)
	fmt.Println(err)
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

func (userD UserDeliveryRealisation) GetCsrf(rwContext echo.Context) error {

	cookie, err := rwContext.Cookie("session_id")
	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	token, _ := csrf.Tokens.Create(cookie.Value, 900+time.Now().Unix())
	csrf := models.Csrf{}
	csrf.Token = token
	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: csrf})
}

func (userD UserDeliveryRealisation) SearchUsers(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	valueOfSearch := rwContext.Param("value")
	valueOfAge := rwContext.QueryParam("year")

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := userD.userLogic.SearchUsers(userId, valueOfSearch, valueOfAge)

	if err != nil {
		userD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.NoContent(http.StatusConflict)
	}

	userD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func NewUserDelivery(log *zap.SugaredLogger, userRealisation users.UserUseCase) UserDeliveryRealisation {
	return UserDeliveryRealisation{userLogic: userRealisation, logger: log}
}

func (userD UserDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/login", userD.Login)
	server.POST("/api/v1/registration", userD.Register)

	server.PUT("/api/v1/settings", userD.UploadSettings)

	server.GET("/api/v1/profile", userD.Profile)
	server.GET("/api/v1/settings", userD.GetSettings)
	server.GET("/api/v1/user/:id", userD.GetUser)
	server.GET("api/v1/users/search/:value", userD.SearchUsers)

	server.GET("api/v1/csrf", userD.GetCsrf)
	server.DELETE("/api/v1/login", userD.Logout)
}
