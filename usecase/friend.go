package usecase

import "github.com/labstack/echo"

type FriendUseCase interface {
	AddFriend(echo.Context, string) error
}
