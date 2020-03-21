package friend

import (
	"../../errors"
	"../../usecase"
	"../../usecase/friend"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type FriendDeliveryRealisation struct {
	friendLogic usecase.FriendUseCase
	logger        *zap.SugaredLogger
}

func (friendD FriendDeliveryRealisation) AddFriend(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	friendD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err := friendD.friendLogic.AddFriend(rwContext , uniqueID.String())

	switch err {
	case errors.InvalidCookie:

		friendD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.NoContent(http.StatusUnauthorized)
	}

	if err != nil {
		friendD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	friendD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewRegisterDelivery(logs *zap.SugaredLogger , friendRealisation friend.FriendUseCaseRealisation) FriendDeliveryRealisation {
	return FriendDeliveryRealisation{friendLogic: friendRealisation, logger: logs}
}
