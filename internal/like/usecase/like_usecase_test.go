package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	lks "main/internal/microservices/likes/delivery"
	err "main/internal/tools/errors"
	mock "main/mocks"
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

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	photoErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.photoId = rand.Int()
		tests.userId = rand.Int()
		lRepoMock.EXPECT().LikePhoto(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.photoId),
		}).Return(nil,photoErr[iter])

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

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	photoErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.photoId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().DislikePhoto(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.photoId),
		}).Return(nil,photoErr[iter])

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

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPostStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().LikePost(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.postId),
		}).Return(nil,postErr[iter])

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

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().DislikePost(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.postId),
		}).Return(nil,postErr[iter])

		errs := likeUseCase.DislikePost(int(tests.postId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}

func TestLikesUseRealisation_DislikeComment(t *testing.T) {

	testLength := 3

	type testPhotoStruct struct {
		postId int
		userId int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().DislikeComment(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.postId),
		}).Return(nil,postErr[iter])

		errs := likeUseCase.DislikeComment(int(tests.postId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}

func TestLikesUseRealisation_LikeComment(t *testing.T) {

	testLength := 3

	type testPhotoStruct struct {
		postId int
		userId int
	}

	ctrl := gomock.NewController(t)

	lRepoMock := mock.NewMockLikeCheckerClient(ctrl)
	likeUseCase := NewLikeUseRealisation(lRepoMock)

	tests := new(testPhotoStruct)
	customErr := errors.New("smth wrong")
	postErr := []error{nil, err.NotExist, customErr}
	expectErr := []error{nil, err.NotExist, customErr}

	for iter := 0; iter < testLength; iter++ {
		tests.postId = rand.Int()
		tests.userId = rand.Int()

		lRepoMock.EXPECT().LikeComment(context.Background(), &lks.Like{
			UserId:               int32(tests.userId),
			DataId:               int32(tests.postId),
		}).Return(nil,postErr[iter])

		errs := likeUseCase.LikeComment(int(tests.postId), int(tests.userId))
		if errs != expectErr[iter] {
			t.Error("Expected value: ", expectErr[iter], " current value: ", errs, " iteration: ", iter)
		}
	}

	ctrl.Finish()
}
