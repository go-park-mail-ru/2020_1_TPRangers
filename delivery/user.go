package delivery

import "github.com/labstack/echo"

type UserDelivery interface {
	GetUser(echo.Context) error
	Profile(echo.Context) error
}
