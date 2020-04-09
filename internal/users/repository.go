package users

import "main/internal/models"

type UserRepository interface {
	GetUserDataByLogin(string) (models.User, error)
	GetUserDataById(int) (models.User, error)
	GetUserProfileSettingsByLogin(string) (models.Settings, error)
	GetUserProfileSettingsById(int) (models.Settings, error)
	UploadSettings(int, models.User) error
	UploadPhoto(string) (int, error)
	GetIdByEmail(string) (int, error)
	GetPassword(string) ([]byte, error)
	GetDefaultProfilePhotoId() (int, error)
	IsUserExist(string) (bool, error)
	GetAlbums(int) ([]models.Album, error)
	GetPhotosFromAlbum(int) (models.Photos, error)
	CreateAlbum(int, models.AlbumReq) error
	UploadPhotoToAlbum(models.PhotoInAlbum) error
	GetUserLoginById(int) (string, error)
	CreateDefaultAlbum(int) error
	AddNewUser(models.User) error
}
