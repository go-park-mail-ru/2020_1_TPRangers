package users

import (
	"main/internal/models"
	"time"
)

type UserUseCase interface {
	GetOtherUserProfileNotLogged(string) (models.OtherUserProfileData, error)
	GetMainUserProfile(int) (models.MainUserProfileData, error)
	GetSettings(int) (models.Settings, error)
	UploadSettings(int, models.Settings) (models.Settings, error)
	Logout(string) error
	Login(models.Auth, string, time.Duration) error
	Register(models.Register, string, time.Duration) error
	GetAlbums(int) ([]models.Album, error)
	GetPhotosFromAlbum(int) (models.Photos, error)
	CreateAlbum(int, models.AlbumReq) error
	UploadPhotoToAlbum(models.PhotoInAlbum) error
	CheckFriendship(int, string) (bool, error)
	GetUserLoginByCookie(int) (string, error)
	GetUserProfileWhileLogged(string, int) (models.OtherUserProfileData, error)

}
