package usecase

import (
	"context"
	lks "main/internal/microservices/likes/delivery"
)

type LikesUseRealisation struct {
	likeMicro lks.LikeCheckerClient
}

func NewLikeUseRealisation(lRepo lks.LikeCheckerClient) LikesUseRealisation {
	return LikesUseRealisation{likeMicro: lRepo}
}

func (Like LikesUseRealisation) LikePhoto(photoId int, userId int) error {
	_, err := Like.likeMicro.LikePhoto(context.Background(), &lks.Like{
		UserId: int32(userId),
		DataId: int32(photoId),
	})
	return err
}

func (Like LikesUseRealisation) DislikePhoto(photoId int, userId int) error {
	_, err := Like.likeMicro.DislikePhoto(context.Background(), &lks.Like{
		UserId: int32(userId),
		DataId: int32(photoId),
	})
	return err
}

func (Like LikesUseRealisation) LikePost(postId int, userId int) error {
	_, err := Like.likeMicro.LikePost(context.Background(), &lks.Like{
		UserId: int32(userId),
		DataId: int32(postId),
	})
	return err
}

func (Like LikesUseRealisation) DislikePost(postId int, userId int) error {
	_, err := Like.likeMicro.DislikePost(context.Background(), &lks.Like{
		UserId: int32(userId),
		DataId: int32(postId),
	})
	return err
}

func (Like LikesUseRealisation) LikeComment(postId int, userId int) error {
	return Like.likeRepo.LikeComment(postId, userId)
}

func (Like LikesUseRealisation) DislikeComment(postId int, userId int) error {
	return Like.likeRepo.DislikeComment(postId, userId)
}
