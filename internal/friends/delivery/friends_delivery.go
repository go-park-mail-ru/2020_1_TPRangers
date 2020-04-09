package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/friends"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"time"
)

type FriendDeliveryRealisation struct {
	friendLogic friends.FriendUseCase
	logger    *zap.SugaredLogger
}


func (userD FriendDeliveryRealisation) FriendList(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	login := rwContext.Param("id")

	friendList, err := userD.friendLogic.GetAllFriends(login)

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

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: friendList})

}

func (userD FriendDeliveryRealisation) GetMainUserFriends(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", err.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	login, err := userD.friendLogic.GetUserLoginByCookie(cookie.Value)
	friendList, err := userD.friendLogic.GetAllFriends(login)

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

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: friendList})

}

func (userD FriendDeliveryRealisation) AddFriend(rwContext echo.Context) error {

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
	err = userD.friendLogic.AddFriend(cookie.Value, friendLogin)

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

func (userD FriendDeliveryRealisation) DeleteFriend(rwContext echo.Context) error {

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
	err = userD.friendLogic.DeleteFriend(cookie.Value, friendLogin)

	errRespStatus := 0

	switch err {
	case errors.InvalidCookie:

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		errRespStatus = http.StatusUnauthorized

	case errors.FailDeleteFriend:
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

func NewUserDelivery(log *zap.SugaredLogger, friendRealisation friends.FriendUseCase) FriendDeliveryRealisation {
	return FriendDeliveryRealisation{friendLogic: friendRealisation, logger: log}
}

func (userD FriendDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.PUT("/api/v1/user/:id", userD.AddFriend)      //

	server.GET("api/v1/friends/:id", userD.FriendList)     //
	server.GET("api/v1/friends", userD.GetMainUserFriends) //

	server.DELETE("/api/v1/user/:id", userD.DeleteFriend)      //
}
