package photos

import (
	"main/internal/models"
)

type PhotoUseCase interface {
	GetPhotosFromAlbum(int) (models.Photos, error)
	UploadPhotoToAlbum(models.PhotoInAlbum) error
}
