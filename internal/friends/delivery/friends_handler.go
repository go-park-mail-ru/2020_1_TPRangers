package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/friends"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
)

type FriendDeliveryRealisation struct {
	friendLogic friends.FriendUseCase
	logger      *zap.SugaredLogger
}

func (userD FriendDeliveryRealisation) FriendList(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

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

	return rwContext.JSON(http.StatusOK, friendList)

}

func (userD FriendDeliveryRealisation) GetMainUserFriends(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	login, err := userD.friendLogic.GetUserLoginById(userId)
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

	return rwContext.JSON(http.StatusOK, friendList)

}

func (userD FriendDeliveryRealisation) AddFriend(rwContext echo.Context) error {

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

	friendLogin := rwContext.Param("id")
	err := userD.friendLogic.AddFriend(userId, friendLogin)

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

	return rwContext.NoContent(http.StatusOK)
}

func (userD FriendDeliveryRealisation) DeleteFriend(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)
	friendLogin := rwContext.Param("id")

	if userId == -1 {

		userD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err := userD.friendLogic.DeleteFriend(userId, friendLogin)

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

	return rwContext.NoContent(http.StatusOK)
}

func (userD FriendDeliveryRealisation) SearchFriends(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	valueOfSearch := rwContext.Param("value")

	if userId == -1 {
		userD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := userD.friendLogic.SearchFriends(userId, valueOfSearch)

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

func NewFriendDelivery(log *zap.SugaredLogger, friendRealisation friends.FriendUseCase) FriendDeliveryRealisation {
	return FriendDeliveryRealisation{friendLogic: friendRealisation, logger: log}
}

func (userD FriendDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/user/:id", userD.AddFriend)
	server.GET("/api/v1/friends/search/:value", userD.SearchFriends)

	server.GET("api/v1/friends/:id", userD.FriendList)
	server.GET("api/v1/friends", userD.GetMainUserFriends)

	server.DELETE("/api/v1/user/:id", userD.DeleteFriend)
}
