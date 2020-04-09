package photos

import (
	"main/internal/models"
)

type PhotoUseCase interface {
	GetPhotosFromAlbum(string, int) (models.Photos, error)
	UploadPhotoToAlbum(string, models.PhotoInAlbum) error
}