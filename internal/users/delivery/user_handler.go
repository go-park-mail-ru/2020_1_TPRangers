package delivery

import (
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"main/internal/csrf"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/internal/users"
	"main/internal/users/usecase"
	"net/http"
	"strconv"
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

	var userData models.OtherUserProfileData
	var err error

	if userId != -1 {
		userData, err = userD.userLogic.GetUserProfileWhileLogged(login, userId)
		userData.IsFriends, err = userD.userLogic.CheckFriendship(userId, login)
	} else {
		userData, err = userD.userLogic.GetOtherUserProfileNotLogged(login)
	}

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

	//cookie, err := rwContext.Cookie("session_id")

	//token := rwContext.Request().Header.Get("X-CSRF-Token")
	//
	//res, err := csrf.Tokens.Check( cookie.Value,  token)
	//fmt.Print(res)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newUserSettings := new(models.Settings)

	err := rwContext.Bind(newUserSettings)

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

	err := rwContext.Bind(userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	info := uuid.NewV4()
	exprTime := 12 * time.Hour
	cookieValue := info.String()

	token, err := userD.userLogic.Login(*userAuthData, cookieValue, exprTime)
	rwContext.Response().Header().Set("X-CSRF-Token", token)

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

	err := rwContext.Bind(userAuthData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	info := uuid.NewV4()
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

func (userD UserDeliveryRealisation) GetAlbums(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	albums, err := userD.userLogic.GetAlbums(userId)

	if err != nil {
		userD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	userD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.JSON(http.StatusOK, albums)
}

func (userD UserDeliveryRealisation) GetPhotosFromAlbum(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	a_id, err := strconv.ParseInt(rwContext.Param("id"), 10, 32)

	photos, err := userD.userLogic.GetPhotosFromAlbum(int(a_id))

	if err != nil {
		userD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	userD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	if len(photos.Urls) == 0 {
		return rwContext.JSON(http.StatusNotFound, photos)
	}

	return rwContext.JSON(http.StatusOK, photos)
}

func (userD UserDeliveryRealisation) CreateAlbum(rwContext echo.Context) error {

	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	albumData := new(models.AlbumReq)

	err := rwContext.Bind(albumData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = userD.userLogic.CreateAlbum(userId, *albumData)

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

	return rwContext.NoContent(http.StatusOK)
}

func (userD UserDeliveryRealisation) UploadPhotoToAlbum(rwContext echo.Context) error {

	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	photoData := new(models.PhotoInAlbum)

	err := rwContext.Bind(photoData)

	if err != nil {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = userD.userLogic.UploadPhotoToAlbum(*photoData)

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

	return rwContext.NoContent(http.StatusOK)

}

func (userD UserDeliveryRealisation) GetCsrf(rwContext echo.Context) error {

	cookie, err := rwContext.Cookie("session_id")
	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	token, _ := csrf.Tokens.Create(cookie.Value,  900 + time.Now().Unix()) // 900 с = 15 минут
	csrf := models.Csrf{}
	csrf.Token = token
	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: csrf})
}

func NewUserDelivery(log *zap.SugaredLogger, userRealisation usecase.UserUseCaseRealisation) UserDeliveryRealisation {
	return UserDeliveryRealisation{userLogic: userRealisation, logger: log}
}

func (userD UserDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/login", userD.Login)
	server.POST("/api/v1/registration", userD.Register)
	server.POST("api/v1/album", userD.CreateAlbum)
	server.POST("api/v1/album/photo", userD.UploadPhotoToAlbum)

	server.PUT("/api/v1/settings", userD.UploadSettings)

	server.GET("/api/v1/profile", userD.Profile)
	server.GET("/api/v1/settings", userD.GetSettings)
	server.GET("/api/v1/user/:id", userD.GetUser)
	server.GET("api/v1/albums", userD.GetAlbums)
	server.GET("api/v1/albums/:id", userD.GetPhotosFromAlbum)
	server.GET("api/v1/csrf", userD.GetCsrf)

	server.DELETE("/api/v1/login", userD.Logout)

}
