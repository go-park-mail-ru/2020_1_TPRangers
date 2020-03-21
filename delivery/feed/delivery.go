package feed

import (
	"../../usecase"
	"../../usecase/feed"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FeedDeliveryRealisation struct {
	feedLogic usecase.FeedUseCase
	logger    *zap.SugaredLogger
}

func (feedD FeedDeliveryRealisation) Feed(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	feedD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err, jsonAnswer := feedD.feedLogic.Feed(rwContext , uniqueID.String())

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusUnauthorized, jsonAnswer)
	}

	feedD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func NewFeedDelivery(log *zap.SugaredLogger, feedRealisation feed.FeedUseCaseRealisation) FeedDeliveryRealisation {
	return FeedDeliveryRealisation{feedLogic: feedRealisation, logger: log}
}
