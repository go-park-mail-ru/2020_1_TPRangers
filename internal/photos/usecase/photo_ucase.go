package usecase

import (
	"context"
	phs "main/internal/microservices/photos/delivery"
	"main/internal/models"
	"main/internal/tools/errors"
)

type PhotoUseCaseRealisation struct {
	photoMicro phs.PhotoCheckerClient
}

func (photoR PhotoUseCaseRealisation) GetPhotosFromAlbum(albumID int) (models.Photos, error) {

	photos, _ := photoR.photoMicro.GetPhotosFromAlbum(context.Background(), &phs.AlbumId{Id: int32(albumID)})

	return models.Photos{
		AlbumName: photos.AlbumName,
		Urls:      photos.Urls,
	}, nil
}

func (photoR PhotoUseCaseRealisation) UploadPhotoToAlbum(photoData models.PhotoInAlbum) error {

	_, err := photoR.photoMicro.UploadPhotoToAlbum(context.Background(), &phs.PhotoInAlbum{
		Url:     photoData.Url,
		AlbumID: photoData.AlbumID,
	})

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func NewPhotoUseCaseRealisation(photoMic phs.PhotoCheckerClient) PhotoUseCaseRealisation {
	return PhotoUseCaseRealisation{
		photoMicro: photoMic,
	}
}
