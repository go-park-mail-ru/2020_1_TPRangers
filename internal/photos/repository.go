package photos

import "main/models"

type PhotoRepository interface {
	GetPhotosFromAlbum(int) (models.Photos, error)
	UploadPhotoToAlbum(models.PhotoInAlbum) error
}
