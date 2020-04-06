package usecase

import(
	"main/internal/like"
	lR "main/internal/like/repository"
	"main/internal/cookies"
	cR "main/internal/cookies/repository"
	"main/internal/tools/errors"
)

type LikesUseRealisation struct {
	likeRepo like.RepositoryLike
	cookieRepo cookies.CookieRepository
}

func NewLikeUseRealisation(lRepo lR.LikeRepositoryRealisation, cRepo cR.CookieRepositoryRealisation ) LikesUseRealisation {
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

func (Like LikesUseRealisation) LikePost(photoId int, cookieValue string) error {
	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.LikePost(photoId, userId)
}

func (Like LikesUseRealisation) DislikePost(photoId int, cookieValue string) error {

	userId , err := Like.cookieRepo.GetUserIdByCookie(cookieValue)

	if err != nil {
		return errors.InvalidCookie
	}

	return Like.likeRepo.DislikePost(photoId, userId)
}





