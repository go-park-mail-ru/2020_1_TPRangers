package delivery

import "github.com/labstack/echo"

type FriendDelivery interface {
	AddFriend(echo.Context) error
}
