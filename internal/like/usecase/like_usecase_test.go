package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	mock "main/mocks"
	err "main/internal/tools/errors"
	"math/rand"
	"testing"
)

func Test_LikePhoto(t *testing.T) {

	testLength := 3

	type testPhotoStruct struct {
		photoId int
		userId  int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	photoErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.photoId = rand.Int()
		tests.userId = rand.Int()
		lRepoMock.EXPECT().LikePhoto(int(tests.photoId), int(tests.userId)).Return(photoErr[iter])

		errs := likeUseCase.LikePhoto(int(tests.photoId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}

func Test_DislikePhoto(t *testing.T) {

	testLength := 3

	type testPhotoStruct struct {
		photoId int
		userId  int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	photoErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.photoId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().DislikePhoto(int(tests.photoId), int(tests.userId)).Return(photoErr[iter])

		errs := likeUseCase.DislikePhoto(int(tests.photoId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}

func Test_LikePost(t *testing.T) {

	testLength := 3

	type testPostStruct struct {
		postId int
		userId int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPostStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().LikePost(int(tests.postId), int(tests.userId)).Return(postErr[iter])

		errs := likeUseCase.LikePost(int(tests.postId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}

func Test_DislikePost(t *testing.T) {

	testLength := 3

	type testPhotoStruct struct {
		postId int
		userId int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().DislikePost(int(tests.postId), int(tests.userId)).Return(postErr[iter])

		errs := likeUseCase.DislikePost(int(tests.postId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}
