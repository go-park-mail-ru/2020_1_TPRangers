package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

type UserUseCase interface {
	GetUser(echo.Context , string) (error, models.JsonStruct)
	Profile(echo.Context , string) (error , models.JsonStruct)
}
