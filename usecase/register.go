package usecase

import "github.com/labstack/echo"

type RegisterUseCase interface {
	Register(echo.Context , string) (error)
}
