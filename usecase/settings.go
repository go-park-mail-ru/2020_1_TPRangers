package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type SettingsUseCase interface {
	GetSettings(echo.Context , string) (error, models.JsonStruct)
	UploadSettings(echo.Context , string) (error, models.JsonStruct)
}
