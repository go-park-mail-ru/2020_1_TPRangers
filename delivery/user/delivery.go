package user

import (
	"../../usecase"
	"../../usecase/user"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserDeliveryRealisation struct {
	userLogic usecase.UserUseCase
	logger    *zap.SugaredLogger
}

func (userD UserDeliveryRealisation) GetUser(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	userD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err, answerJson := userD.userLogic.GetUser(rwContext, uniqueID.String())

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusNotFound),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusNotFound, answerJson)
	}

	userD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.JSON(http.StatusOK, answerJson)
}

func (userD UserDeliveryRealisation) Profile(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	userD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err, answerJson := userD.userLogic.Profile(rwContext, uniqueID.String())

	if err != nil {

		userD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusUnauthorized, answerJson)
	}

	userD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.JSON(http.StatusOK, answerJson)
}

func NewUserDelivery(log *zap.SugaredLogger, userRealisation user.UserUseCaseRealisation) UserDeliveryRealisation {
	return UserDeliveryRealisation{userLogic: userRealisation, logger: log}
}
