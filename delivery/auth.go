package delivery

import "github.com/labstack/echo"

type AuthDelivery interface {
	Login(echo.Context) error
	Logout(echo.Context) error
}