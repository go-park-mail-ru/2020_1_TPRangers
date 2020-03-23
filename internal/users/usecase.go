package users

import (
	"../models"
	"time"
)

type UserUseCase interface {
	GetUser(string) (map[string]interface{}, error)
	Profile(string) (map[string]interface{}, error)
	GetSettings(string) (map[string]interface{}, error)
	UploadSettings(string, models.Settings) (map[string]interface{}, error)
	Logout(string) error
	Login(models.Auth, string, time.Duration) error
	AddFriend(string, string) error
	Register(models.Register, string, time.Duration) error
}