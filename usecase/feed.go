package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type FeedUseCase interface {
	Feed(echo.Context , string) (error, models.JsonStruct)
}
