package delivery

import "github.com/labstack/echo"

type RegisterDelivery interface {
	Register(echo.Context) error
}
