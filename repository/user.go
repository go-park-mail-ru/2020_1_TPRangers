package repository

import "../models"

type UserRepository interface {
	GetUserDataByLogin(string) (models.User, error)
	GetUserDataById(int) (models.User, error)
	UploadSettings(int, models.User) error
}
