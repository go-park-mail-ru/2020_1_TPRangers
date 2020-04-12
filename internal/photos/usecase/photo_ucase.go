package usecase

import (
	"main/internal/models"
	"main/internal/photos"
	photoRep "main/internal/photos/repository"
	"main/internal/tools/errors"
)

type PhotoUseCaseRealisation struct {
	photoDB photos.PhotoRepository
}

func (photoR PhotoUseCaseRealisation) GetPhotosFromAlbum(albumID int) (models.Photos, error) {

	photos, _ := photoR.photoDB.GetPhotosFromAlbum(albumID)

	return photos, nil
}

func (photoR PhotoUseCaseRealisation) UploadPhotoToAlbum(photoData models.PhotoInAlbum) error {

	err := photoR.photoDB.UploadPhotoToAlbum(photoData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func NewPhotoUseCaseRealisation(photoDB photoRep.PhotoRepositoryRealisation) PhotoUseCaseRealisation {
	return PhotoUseCaseRealisation{
		photoDB: photoDB,
	}
}
