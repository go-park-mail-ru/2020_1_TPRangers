package auth

import (
	"../../errors"
	"../../usecase"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LoginDeliveryRealisation struct {
	loginLogic usecase.AuthUseCase
	logger     *zap.SugaredLogger
}

func (logD LoginDeliveryRealisation) Login(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	logD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err := logD.loginLogic.Login(rwContext, uniqueID.String())

	switch err {
	case errors.WrongLogin, errors.WrongPassword:

		logD.logger.Info(
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
		logD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", "convertion error"),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)
		return rwContext.NoContent(http.StatusInternalServerError)
	}

	logD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.NoContent(http.StatusOK)
}

func (logD LoginDeliveryRealisation) Logout(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	logD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	logD.loginLogic.Logout(rwContext, uniqueID.String())

	logD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewLoginDelivery(log *zap.SugaredLogger , lLogic usecase.AuthUseCase) LoginDeliveryRealisation {
	return LoginDeliveryRealisation{loginLogic: lLogic, logger: log}
}
