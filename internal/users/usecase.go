package users

import (
	"main/internal/models"
)

type UserUseCase interface {
	GetOtherUserProfileNotLogged(string) (models.OtherUserProfileData, error)
	GetMainUserProfile(int) (models.MainUserProfileData, error)
	GetSettings(int) (models.Settings, error)
	UploadSettings(int, models.Settings) (models.Settings, error)
	Logout(string) error
	Login(userData models.Auth) (string , error)
	Register(models.Register) (string , error)
	CheckFriendship(int, string) (bool, error)
	GetUserLoginByCookie(int) (string, error)
	GetUserProfileWhileLogged(string, int) (models.OtherUserProfileData, error)
	SearchUsers(int, string) ([]models.Person, error)
}
