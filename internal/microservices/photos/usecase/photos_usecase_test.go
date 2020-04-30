package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"main/internal/microservices/photos/delivery"
	"main/internal/models"
	errors2 "main/internal/tools/errors"
	"main/mocks"
	"testing"
)

func TestPhotoUseChecker_GetPhotosFromAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	pDBMock := mock.NewMockPhotoRepository(ctrl)
	pTest := NewPhotoUseCaseChecker(pDBMock)

	id := 1
	newId := &photos.AlbumId{
		Id:                   int32(id),
	}

	pDBMock.EXPECT().GetPhotosFromAlbum(id).Return(models.Photos{
		AlbumName: "123",
		Urls:      nil,
	}, nil)

	if data , errs := pTest.GetPhotosFromAlbum(context.Background(),newId); data.AlbumName != "123" || data.Urls != nil || errs != nil {
		t.Error("ERROR")
	}
}

func TestPhotoUseChecker_UploadPhotoToAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	pDBMock := mock.NewMockPhotoRepository(ctrl)
	pTest := NewPhotoUseCaseChecker(pDBMock)

	url := "123"
	albumId := "123"
	phInAlb := &photos.PhotoInAlbum{
		Url:                  url,
		AlbumID:              albumId,
	}

	pDBMock.EXPECT().UploadPhotoToAlbum(models.PhotoInAlbum{
		Url:     url,
		AlbumID: albumId,
	}).Return(nil)

	if _ , errs := pTest.UploadPhotoToAlbum(context.Background(), phInAlb); errs != nil {
		t.Error("ERROR")
	}

	pDBMock.EXPECT().UploadPhotoToAlbum(models.PhotoInAlbum{
		Url:     url,
		AlbumID: albumId,
	}).Return(errors.New("123"))

	if _ , errs := pTest.UploadPhotoToAlbum(context.Background(), phInAlb); errs != errors2.FailReadFromDB{
		t.Error("ERROR")
	}
}
