package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type FeedUseCase interface {
	Feed(ctx echo.Context) (error, models.JsonStruct)
}
