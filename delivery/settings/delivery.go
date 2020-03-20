package settings

import (
	"../../errors"
	"../../usecase"
	"../../usecase/settings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type SettingsDeliveryRealisation struct {
	settingsLogic usecase.SettingsUseCase
	logger        *zap.SugaredLogger
}

func (setD SettingsDeliveryRealisation) GetSettings(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	setD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err, jsonAnswer := setD.settingsLogic.GetSettings(rwContext, uniqueID.String())

	switch err {
	case errors.InvalidCookie:
		setD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusUnauthorized, jsonAnswer)

	case errors.FailReadFromDB:
		setD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusInternalServerError, jsonAnswer)
	}

	if err != nil {

		setD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusInternalServerError, jsonAnswer)

	}

	setD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func (setD SettingsDeliveryRealisation) UploadSettings(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	setD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
	)

	err, jsonAnswer := setD.settingsLogic.UploadSettings(rwContext, uniqueID.String())

	switch err {
	case errors.InvalidCookie:
		setD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusUnauthorized, jsonAnswer)

	case errors.FailReadFromDB:
		setD.logger.Info(
			zap.String("ID", uniqueID.String()),
			zap.String("URL", rwContext.Request().URL.Path),
			zap.String("METHOD", rwContext.Request().Method),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			zap.Duration("TIME FOR ANSWER", time.Since(start)),
		)

		return rwContext.JSON(http.StatusInternalServerError, jsonAnswer)
	}

	setD.logger.Info(
		zap.String("ID", uniqueID.String()),
		zap.String("URL", rwContext.Request().URL.Path),
		zap.String("METHOD", rwContext.Request().Method),
		zap.Int("ANSWER STATUS", http.StatusOK),
		zap.Duration("TIME FOR ANSWER", time.Since(start)),
	)

	return rwContext.JSON(http.StatusOK, jsonAnswer)

}

func NewSettingsDelivery(log *zap.SugaredLogger, setRealisation settings.SettingsUseCaseRealisation) SettingsDeliveryRealisation {
	return SettingsDeliveryRealisation{settingsLogic: setRealisation, logger: log}
}
