package usecase

import "github.com/labstack/echo"

type AuthUseCase interface {
	Logout(echo.Context) error
	Login(echo.Context) (error, string)
}
