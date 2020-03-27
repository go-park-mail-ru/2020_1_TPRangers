package users

import "main/internal/models"

type UserRepository interface {
	GetUserDataByLogin(string) (models.User, error)
	GetUserDataById(int) (models.User, error)
	GetUserProfileSettingsByLogin(string) (models.Settings , error)
	GetUserProfileSettingsById(int) (models.Settings , error)
	UploadSettings(int, models.User) error
	UploadPhoto(string) (int , error)
	GetIdByEmail(string) (int, error)
	GetPassword(string) (string, error)
	AddNewUser(models.User) error
	GetDefaultProfilePhotoId() (int , error)
	IsUserExist(string) (bool, error)
	AddFriend(int, int) error
	GetFriendIdByLogin(string) (int, error)
}