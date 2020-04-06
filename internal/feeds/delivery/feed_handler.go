package delivery

import (
	"main/internal/tools/errors"
	"main/internal/models"
	"main/internal/feeds"
	"main/internal/feeds/usecase"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FeedDeliveryRealisation struct {
	feedLogic feeds.FeedUseCase
	logger    *zap.SugaredLogger
}

func (feedD FeedDeliveryRealisation) Feed(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := feedD.feedLogic.Feed(cookie.Value)

	errRespStauts := 0

	switch err {
	case errors.InvalidCookie:

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		errRespStauts = http.StatusUnauthorized
	case errors.FailReadFromDB:
		errRespStauts = http.StatusInternalServerError
	}

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStauts),
		)

		return rwContext.JSON(errRespStauts, models.JsonStruct{Err: err.Error()})
	}

	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: jsonAnswer})
}

func (feedD FeedDeliveryRealisation) CreatePost(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newPost := new(models.Post)

	err = rwContext.Bind(&newPost)

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}


	err = feedD.feedLogic.CreatePost(cookie.Value, *newPost)

	errRespStauts := 0

	switch err {
	case errors.InvalidCookie:
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		errRespStauts = http.StatusUnauthorized
	case errors.FailSendToDB:
		errRespStauts = http.StatusInternalServerError
	}

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStauts),
		)

		return rwContext.JSON(errRespStauts, models.JsonStruct{Err: err.Error()})
	}

	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewFeedDelivery(log *zap.SugaredLogger, feedRealisation usecase.FeedUseCaseRealisation) FeedDeliveryRealisation {
	return FeedDeliveryRealisation{feedLogic: feedRealisation, logger: log}
}

func (feedD FeedDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/news", feedD.Feed)
	server.POST("/api/v1/post", feedD.CreatePost)

}
