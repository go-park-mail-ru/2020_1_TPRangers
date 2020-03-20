package repository

import "../models"

type RegisterRepository interface {
	AddNewUser(models.User) error
	GetDefaultProfilePhotoId() (int , error)
	GetIdByEmail(string) (int, error)
	IsUserExist(string) (bool, error)
}
