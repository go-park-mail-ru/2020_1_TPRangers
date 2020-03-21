package usecase

import "github.com/labstack/echo"

type AuthUseCase interface {
	Logout(echo.Context , string) error
	Login(echo.Context , string) error
}
