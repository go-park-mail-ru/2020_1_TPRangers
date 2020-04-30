package usecase

import (
	"github.com/golang/mock/gomock"
	"main/internal/tools/errors"
	"main/mocks"
	"main/models"
	"math/rand"
	"testing"
)

func TestAlbumUseCaseRealisation_GetAlbums(t *testing.T) {

	ctrl := gomock.NewController(t)

	aRepoMock := mock.NewMockAlbumRepository(ctrl)
	albumUse := NewAlbumUseCaseRealisation(aRepoMock)

	errs := []error{nil, errors.FailSendToDB}
	expectBehaviour := []error{nil, errors.FailSendToDB}

	for iter, _ := range expectBehaviour {

		uId := rand.Int()
		albums := make([]models.Album, 1, 1)
		if errs[iter] != nil {
			albums[0] = models.Album{}
		} else {
			albums[0] = models.Album{
				ID:       "123",
				Name:     "222",
				PhotoUrl: nil,
			}
		}

		aRepoMock.EXPECT().GetAlbums(uId).Return(albums, errs[iter])

		if alb, err := albumUse.GetAlbums(uId); err != expectBehaviour[iter] || alb[0] != albums[0] {
			t.Error(iter, err, expectBehaviour[iter])
		}

	}

}

func TestAlbumUseCaseRealisation_CreateAlbum(t *testing.T) {

	ctrl := gomock.NewController(t)

	aRepoMock := mock.NewMockAlbumRepository(ctrl)
	albumUse := NewAlbumUseCaseRealisation(aRepoMock)

	errs := []error{nil, errors.FailSendToDB}
	expectBehaviour := []error{nil, errors.FailReadFromDB}

	for iter, _ := range expectBehaviour {

		uId := rand.Int()

		albumData := new(models.AlbumReq)

		aRepoMock.EXPECT().CreateAlbum(uId, *albumData).Return(errs[iter])

		if err := albumUse.CreateAlbum(uId, *albumData); err != expectBehaviour[iter] {
			t.Error(iter, err, expectBehaviour[iter])
		}

	}

}
