package repository

import "../models"

type UserRepository interface {
	GetUserDataByLogin(string) (models.User, error)
	GetUserDataById(int) (models.User, error)
	GetUserProfileSettingsByLogin(string) (models.Settings , error)
	GetUserProfileSettingsById(int) (models.Settings , error)
	UploadSettings(int, models.User) error
	UploadPhoto(string) (int , error)
}
