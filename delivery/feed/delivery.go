package feed

import (
	"../../errors"
	"../../models"
	"../../usecase"
	"../../usecase/feed"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FeedDeliveryRealisation struct {
	feedLogic usecase.FeedUseCase
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

func NewFeedDelivery(log *zap.SugaredLogger, feedRealisation feed.FeedUseCaseRealisation) FeedDeliveryRealisation {
	return FeedDeliveryRealisation{feedLogic: feedRealisation, logger: log}
}

func (feedD FeedDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/news", feedD.Feed)
}
