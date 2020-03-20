package usecase

import (
	"../models"
	"github.com/labstack/echo"
)

func GetDataFromJson(jsonType string, r echo.Context) (data interface{}, errConvert error) {

	switch jsonType {

	case "reg":
		data = new(models.Register)
	case "data":
		data = new(models.User)
	case "log":
		data = new(models.Auth)
	}

	errConvert = r.Bind(data)
	return
}
