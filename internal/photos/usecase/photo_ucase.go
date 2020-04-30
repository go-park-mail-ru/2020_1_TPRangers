package usecase

import (
	"context"
	phs "main/internal/microservices/photos/delivery"
	"main/internal/models"
	"main/internal/tools/errors"
	"fmt"
	"os"

)

type PhotoUseCaseRealisation struct {
	photoMicro phs.PhotoCheckerClient
}

func (photoR PhotoUseCaseRealisation) GetPhotosFromAlbum(albumID int) (models.Photos, error) {

	photos, _ := photoR.photoMicro.GetPhotosFromAlbum(context.Background(), &phs.AlbumId{Id: int32(albumID)})

	file, err := os.Create("hello.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("Album name is")
	file.WriteString(photos.AlbumName)

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
