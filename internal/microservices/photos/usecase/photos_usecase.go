package usecase

import (
	"context"
	"main/internal/microservices/photos/delivery"
	"main/internal/models"
	phs "main/internal/photos"
	photoRep "main/internal/photos/repository"
	"main/internal/tools/errors"
)

type PhotoUseChecker struct {
	photoDB phs.PhotoRepository
}

func (photoR PhotoUseChecker) GetPhotosFromAlbum(ctx context.Context, id *photos.AlbumId) (*photos.Photos, error) {

	phs, _ := photoR.photoDB.GetPhotosFromAlbum(int(id.Id))

	return &photos.Photos{
		AlbumName: phs.AlbumName,
		Urls:      phs.Urls,
	}, nil
}

func (photoR PhotoUseChecker) UploadPhotoToAlbum(ctx context.Context, newPhoto *photos.PhotoInAlbum) (*photos.Dummy, error) {

	err := photoR.photoDB.UploadPhotoToAlbum(models.PhotoInAlbum{
		Url:     newPhoto.Url,
		AlbumID: newPhoto.AlbumID,
	})

	if err != nil {
		return &photos.Dummy{}, errors.FailReadFromDB
	}

	return &photos.Dummy{}, nil
}

func NewPhotoUseCaseChecker(photoDB photoRep.PhotoRepositoryRealisation) PhotoUseChecker {
	return PhotoUseChecker{
		photoDB: photoDB,
	}
}
