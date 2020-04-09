package usecase

import (
	"main/internal/cookies"
	"main/internal/like"
	"main/internal/tools/errors"
)

type LikesUseRealisation struct {
	likeRepo like.RepositoryLike
	cookieRepo cookies.CookieRepository
}

func NewLikeUseRealisation(lRepo like.RepositoryLike, cRepo cookies.CookieRepository) LikesUseRealisation {
	return LikesUseRealisation{likeRepo:lRepo, cookieRepo:cRepo}
}

func (Like LikesUseRealisation) LikePhoto(photoId int, cookieValue string) error {
	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.LikePhoto(photoId, userId)
}

func (Like LikesUseRealisation) DislikePhoto(photoId int, cookieValue string) error {

	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.DislikePhoto(photoId, userId)
}

func (Like LikesUseRealisation) LikePost(postId int, cookieValue string) error {
	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.LikePost(postId, userId)
}

func (Like LikesUseRealisation) DislikePost(postId int, cookieValue string) error {

	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.DislikePost(postId, userId)
}





