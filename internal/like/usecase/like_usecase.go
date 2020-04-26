package usecase

import (
	"main/internal/like"
)

type LikesUseRealisation struct {
	likeRepo like.RepositoryLike
}

func NewLikeUseRealisation(lRepo like.RepositoryLike) LikesUseRealisation {
	return LikesUseRealisation{likeRepo: lRepo}
}

func (Like LikesUseRealisation) LikePhoto(photoId int, userId int) error {
	return Like.likeRepo.LikePhoto(photoId, userId)
}

func (Like LikesUseRealisation) DislikePhoto(photoId int, userId int) error {
	return Like.likeRepo.DislikePhoto(photoId, userId)
}

func (Like LikesUseRealisation) LikePost(postId int, userId int) error {
	return Like.likeRepo.LikePost(postId, userId)
}

func (Like LikesUseRealisation) DislikePost(postId int, userId int) error {
	return Like.likeRepo.DislikePost(postId, userId)
}

func (Like LikesUseRealisation) LikeComment(postId int, userId int) error {
	return Like.likeRepo.LikeComment(postId, userId)
}

func (Like LikesUseRealisation) DislikeComment(postId int, userId int) error {
	return Like.likeRepo.DislikeComment(postId, userId)
}
