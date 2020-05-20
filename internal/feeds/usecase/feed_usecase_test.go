package usecase

import (
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/mocks"
	"testing"
)

func TestFeedUseCaseRealisation_Feed(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock)

	feedErr := []error{nil, errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB}
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
	}}, nil}

	for iter, _ := range expectValues {
		userId := 30
		fRepoMock.EXPECT().GetUserFeedById(userId, 30).Return(expectValues[iter], feedErr[iter])

		if _, err := fRepoTest.Feed(userId); err != expectErr[iter] {
			t.Error("unexpected behaviour")
		}
	}

}

func TestFeedUseCaseRealisation_CreatePost(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock)

	feedErr := []error{nil, errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB}


	for iter, _ := range expectErr {
		userId := 30
		userLogin := uuid.NewV4()
		post := models.Post{
			Id:            0,
			Text:          "123123123123123123",
		}
		fRepoMock.EXPECT().CreatePost(userId, userLogin.String(),post).Return(feedErr[iter])

		if err := fRepoTest.CreatePost(userId, userLogin.String(),post); err != expectErr[iter] {
			t.Error("unexpected behaviour")
		}
	}

}

func TestFeedUseCaseRealisation_CreateComment(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock)

	feedErr := []error{nil, errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB}


	for iter, _ := range expectErr {
		userId := 30
		userLogin := uuid.NewV4()
		comment := models.Comment{
			Text:          userLogin.String(),
		}
		fRepoMock.EXPECT().CreateComment(userId, comment).Return(feedErr[iter])

		if err := fRepoTest.CreateComment(userId, comment); err != expectErr[iter] {
			t.Error("unexpected behaviour")
		}
	}

}

func TestFeedUseCaseRealisation_DeleteComment(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock)

	feedErr := []error{nil, errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB}


	for iter, _ := range expectErr {
		userId := 30
		commentId := uuid.NewV4()

		fRepoMock.EXPECT().DeleteComment(userId, commentId.String()).Return(feedErr[iter])

		if err := fRepoTest.DeleteComment(userId, commentId.String()); err != expectErr[iter] {
			t.Error("unexpected behaviour")
		}
	}

}

func TestFeedUseCaseRealisation_GetPostAndComments(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fRepoMock := mock.NewMockFeedRepository(ctrl)
	fRepoTest := NewFeedUseCaseRealisation(fRepoMock)

	feedErr := []error{nil, errors.FailReadFromDB}
	expectErr := []error{nil, errors.FailReadFromDB}


	for iter, _ := range expectErr {
		userId := 30
		commentId := uuid.NewV4()
		post := models.Post{
			Id:            0,
			Text:          "123123123123123123",
		}

		fRepoMock.EXPECT().GetPostAndComments(userId, commentId.String()).Return(post,feedErr[iter])

		if _ , err := fRepoTest.GetPostAndComments(userId, commentId.String()); err != expectErr[iter] {
			t.Error("unexpected behaviour")
		}
	}

}
