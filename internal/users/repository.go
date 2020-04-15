package users

import "main/internal/models"

type UserRepository interface {
	GetIdByLogin(string) (int, error)
	GetUserDataByLogin(string) (models.User, error)
	GetUserDataById(int) (models.User, error)
	GetUserProfileSettingsByLogin(string) (models.Settings, error)
	GetUserProfileSettingsById(int) (models.Settings, error)
	UploadSettings(int, models.User) error
	UploadProfilePhoto(string) (int, error)
	GetIdByEmail(string) (int, error)
	GetPassword(string) ([]byte, error)
	GetDefaultProfilePhotoId() (int, error)
	IsUserExist(string) (bool, error)
	GetUserLoginById(int) (string, error)
	AddNewUser(models.User) error
}
