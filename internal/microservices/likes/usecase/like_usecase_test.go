package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	lproto "main/internal/microservices/likes/delivery"
	"main/mocks"
	"testing"
)

func TestLikesUseChecker_LikePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().LikePhoto(dataId,userId).Return(customErr)

	if _ , err := likeTest.LikePhoto(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.LikePhoto(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}

func TestLikesUseChecker_LikePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().LikePost(dataId,userId).Return(customErr)

	if _ , err := likeTest.LikePost(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.LikePost(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}

func TestLikesUseChecker_LikeComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().LikeComment(dataId,userId).Return(customErr)

	if _ , err := likeTest.LikeComment(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.LikeComment(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}

func TestLikesUseChecker_DislikePhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().DislikePhoto(dataId,userId).Return(customErr)

	if _ , err := likeTest.DislikePhoto(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.DislikePhoto(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}

func TestLikesUseChecker_DislikePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().DislikePost(dataId,userId).Return(customErr)

	if _ , err := likeTest.DislikePost(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.DislikePost(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}

func TestLikesUseChecker_DislikeComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	lRepoMock := mock.NewMockRepositoryLike(ctrl)
	likeTest := NewLikeUseCaseChecker(lRepoMock)

	customErr := errors.New("123")
	dataId := 1
	userId := 1
	likeData := &lproto.Like{
		UserId:               int32(userId),
		DataId:               int32(dataId),
	}
	lRepoMock.EXPECT().DislikeComment(dataId,userId).Return(customErr)

	if _ , err := likeTest.DislikeComment(context.Background(),likeData); err != customErr{
		t.Error("ERROR", err)
	}

	if _ , err := likeTest.DislikeComment(context.Background(),nil); err != nil{
		t.Error("ERROR",err)
	}

}