package usecase

import (
	"context"
	"main/internal/like"
	lproto "main/internal/microservices/likes/delivery"
)

type LikesUseChecker struct {
	likeRepo like.RepositoryLike
}

func NewLikeUseCaseChecker(lRepo like.RepositoryLike) LikesUseChecker {
	return LikesUseChecker{likeRepo: lRepo}
}

func (Like LikesUseChecker) LikePhoto(ctx context.Context, lPhoto *lproto.Like) (*lproto.Dummy, error) {

	if lPhoto == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.LikePhoto(int(lPhoto.DataId), int(lPhoto.UserId))
}

func (Like LikesUseChecker) DislikePhoto(ctx context.Context, disPhoto *lproto.Like) (*lproto.Dummy, error) {

	if disPhoto == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.DislikePhoto(int(disPhoto.DataId), int(disPhoto.UserId))
}

func (Like LikesUseChecker) LikePost(ctx context.Context, lPost *lproto.Like) (*lproto.Dummy, error) {

	if lPost == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.LikePost(int(lPost.DataId), int(lPost.UserId))
}

func (Like LikesUseChecker) DislikePost(ctx context.Context, disPost *lproto.Like) (*lproto.Dummy, error) {

	if disPost == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.DislikePost(int(disPost.DataId), int(disPost.UserId))
}

func (Like LikesUseChecker) LikeComment(ctx context.Context, lComm *lproto.Like) (*lproto.Dummy, error) {

	if lComm == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.LikeComment(int(lComm.DataId), int(lComm.UserId))
}

func (Like LikesUseChecker) DislikeComment(ctx context.Context, disComm *lproto.Like) (*lproto.Dummy, error) {

	if disComm == nil {
		return &lproto.Dummy{}, nil
	}

	return &lproto.Dummy{}, Like.likeRepo.DislikeComment(int(disComm.DataId), int(disComm.UserId))
}
