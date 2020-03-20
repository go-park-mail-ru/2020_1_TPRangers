package delivery

import "github.com/labstack/echo"

type FeedDelivery interface {
	Feed(echo.Context) error
}
