package register

import (
	"../../errors"
	"../../usecase"
	"../../usecase/register"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type RegisterDeliveryRealisation struct {
	registerLogic usecase.RegisterUseCase
	logger        *zap.SugaredLogger
}

func (regD RegisterDeliveryRealisation) Register(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	regD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err := regD.registerLogic.Register(rwContext , uniqueID.String())

	switch err {
	case errors.AlreadyExist:

		regD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	if err != nil {
		regD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", "convertion error"),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	regD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewRegisterDelivery(logs *zap.SugaredLogger , regRealisation register.RegisterUseCaseRealisation) RegisterDeliveryRealisation {
	return RegisterDeliveryRealisation{registerLogic: regRealisation, logger: logs}
}
