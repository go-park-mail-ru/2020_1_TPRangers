package users

import (
	"main/internal/models"
	"time"
)

type UserUseCase interface {
	GetUser(string) (map[string]interface{}, error)
	Profile(string) (map[string]interface{}, error)
	GetSettings(string) (map[string]interface{}, error)
	UploadSettings(string, models.Settings) (map[string]interface{}, error)
	Logout(string) error
	Login(models.Auth, string, time.Duration) error
	Register(models.Register, string, time.Duration) error
	GetAlbums(string) ([]models.Album, error)
	GetPhotosFromAlbum(string, int) (models.Photos, error)
	UploadPhotoToAlbum(string, models.PhotoInAlbum) error
	CreateAlbum(string, models.AlbumReq) error

	CheckFriendship(string, string, map[string]interface{}) (map[string]interface{}, error)
	GetUserLoginByCookie(string) (string, error)
	GetUserWhileLogged(string,string) (map[string]interface{}, error)
}

