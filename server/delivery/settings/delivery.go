package settings

import (
	"../../usecase"
	"../../usecase/settings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type SettingsDeliveryRealisation struct {
	settingsLogic usecase.SettingsUseCase
}

func (setD SettingsDeliveryRealisation) GetSettings(rwContext echo.Context) error {
	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, jsonAnswer := setD.settingsLogic.GetSettings(rwContext)

	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized, jsonAnswer)
	}

	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func (setD SettingsDeliveryRealisation) UploadSettings(rwContext echo.Context) error {

	uniqueID, _ := uuid.NewV4()
	//start := time.Now()
	rwContext.Response().Header().Set("REQUEST_ID", uniqueID.String())

	err, jsonAnswer := setD.settingsLogic.UploadSettings(rwContext)

	if err != nil {
		return rwContext.JSON(http.StatusUnauthorized, jsonAnswer)
	}

	return rwContext.JSON(http.StatusOK, jsonAnswer)

}

func NewSettingsDelivery() SettingsDeliveryRealisation {
	settingsHandler := settings.SettingsUseCaseRealisation{}
	return SettingsDeliveryRealisation{settingsLogic: settingsHandler}
}
