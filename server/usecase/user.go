package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type UserUseCase interface {
	GetUser(echo.Context) (error, models.JsonStruct)
}
