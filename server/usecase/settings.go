package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type SettingsUseCase interface {
	GetSettings(echo.Context) (error, models.JsonStruct)
	UploadSettings(echo.Context) (error, models.JsonStruct)
}
