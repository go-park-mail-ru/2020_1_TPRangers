package photos

import (
	"main/models"
)

type PhotoUseCase interface {
	GetPhotosFromAlbum(int) (models.Photos, error)
	UploadPhotoToAlbum(models.PhotoInAlbum) error
}
