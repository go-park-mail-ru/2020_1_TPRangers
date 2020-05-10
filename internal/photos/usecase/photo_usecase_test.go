package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	phss "main/internal/microservices/photos/delivery"
	"main/internal/models"
	errs "main/internal/tools/errors"
	"main/mocks"
	"testing"
)

func TestPhotoUseCaseRealisation_GetPhotosFromAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	photoMicro := mock.NewMockPhotoCheckerClient(ctrl)
	photoTest := NewPhotoUseCaseRealisation(photoMicro)

	customErr := errors.New("smth happend")
	albumId := 1
	phs := models.Photos{
		AlbumName: "fuck",
		Urls:      []string{"xd", "fuck"},
	}

	grpcPhs := phss.Photos{
		AlbumName: "fuck",
		Urls:      []string{"xd", "fuck"},
	}

	photoMicro.EXPECT().GetPhotosFromAlbum(context.Background(), &phss.AlbumId{
		Id: int32(albumId),
	}).Return(&grpcPhs, customErr)

	if ps, err := photoTest.GetPhotosFromAlbum(albumId); ps.AlbumName != phs.AlbumName || err != nil {
		t.Error("ERROR")
	}

}

func TestPhotoUseCaseRealisation_UploadPhotoToAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	photoMicro := mock.NewMockPhotoCheckerClient(ctrl)
	photoTest := NewPhotoUseCaseRealisation(photoMicro)

	customErr := errors.New("smth happend")
	phs := models.PhotoInAlbum{
		Url:     "fuck",
		AlbumID: "2",
	}

	grpcPhs := phss.PhotoInAlbum{
		Url:     "fuck",
		AlbumID: "2",
	}

	photoMicro.EXPECT().UploadPhotoToAlbum(context.Background(), &grpcPhs).Return(nil, nil)
	if err := photoTest.UploadPhotoToAlbum(phs); err != nil {
		fmt.Println(err)
		t.Error("ERROR")
	}

	photoMicro.EXPECT().UploadPhotoToAlbum(context.Background(), &grpcPhs).Return(nil, customErr)

	if err := photoTest.UploadPhotoToAlbum(phs); err != errs.FailReadFromDB {
		t.Error("ERROR")
	}

}
