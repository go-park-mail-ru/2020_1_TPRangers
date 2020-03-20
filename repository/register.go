package repository

import "../models"

type RegisterRepository interface {
	AddNewUser(models.User) error
	GetIdByEmail(string) (int, error)
	IsUserExist(string) (bool, error)
}
