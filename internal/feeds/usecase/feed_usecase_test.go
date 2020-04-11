package usecase

import (
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"main/internal/feeds/usecase/mock"
	"main/internal/models"
	"main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestFeedUseCaseRealisation_Feed(t *testing.T) {

	cVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cRepoMock := mock.NewMockCookieRepository(ctrl)
	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock,cRepoMock)

	cookieErr := []error{nil, nil, errors.InvalidCookie , errors.InvalidCookie}
	feedErr := []error{nil , errors.FailReadFromDB , nil , errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB , errors.InvalidCookie , errors.InvalidCookie}
	expectValues := [][]models.Post{[]models.Post{models.Post{
		Id:            0,
		Text:          "",
		Photo:         models.Photo{},
		Attachments:   "",
		Likes:         0,
		WasLike:       false,
		Creation:      "",
		AuthorName:    "",
		AuthorSurname: "",
		AuthorUrl:     "",
		AuthorPhoto:   "",
	}}, nil , nil , nil}

	for iter , _ := range(expectValues) {

		uId := rand.Int()
		cookieVal := cVal.String()

		cRepoMock.EXPECT().GetUserIdByCookie(cookieVal).Return(uId , cookieErr[iter])
		if cookieErr[iter] == nil {
			fRepoMock.EXPECT().GetUserFeedById(uId,30).Return(expectValues[iter] , feedErr[iter])
		}

		eVal , eErr := fRepoTest.Feed(cookieVal)

		if eErr != expectErr[iter] {
			t.Error("expected value :" , expectValues[iter] , " got value : ", eVal)
		}


	}

}

func TestFeedUseCaseRealisation_CreatePost(t *testing.T) {

	cVal := uuid.NewV4()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cRepoMock := mock.NewMockCookieRepository(ctrl)
	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock,cRepoMock)

	cookieErr := []error{nil, nil, errors.InvalidCookie , errors.InvalidCookie}
	createErr := []error{nil , errors.FailReadFromDB , nil , errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB , errors.InvalidCookie , errors.InvalidCookie}
	expectValues := models.Post{
		Id:            0,
		Text:          "",
		Photo:         models.Photo{},
		Attachments:   "",
		Likes:         0,
		WasLike:       false,
		Creation:      "",
		AuthorName:    "",
		AuthorSurname: "",
		AuthorUrl:     "",
		AuthorPhoto:   "",
	}

	for iter , _ := range(expectErr) {

		uId := rand.Int()
		cookieVal := cVal.String()

		cRepoMock.EXPECT().GetUserIdByCookie(cookieVal).Return(uId , cookieErr[iter])
		if cookieErr[iter] == nil {
			fRepoMock.EXPECT().CreatePost(uId,expectValues).Return(createErr[iter])
		}

		eErr := fRepoTest.CreatePost(cookieVal, expectValues)

		if eErr != expectErr[iter] {
			t.Error("expected value :" , expectErr[iter] , " got value : ", eErr)
		}


	}


}
