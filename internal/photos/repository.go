package photos

import "main/internal/models"

type PhotoRepository interface {
	GetPhotosFromAlbum(int) (models.Photos, error)
	UploadPhotoToAlbum(models.PhotoInAlbum) error
}
