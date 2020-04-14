package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/like"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)

type LikeDelivery struct {
	likeLogic like.UseCaseLike
	logger    *zap.SugaredLogger
}

func NewLikeDelivery(log *zap.SugaredLogger, likeRealisation like.UseCaseLike) LikeDelivery {
	return LikeDelivery{likeLogic: likeRealisation, logger: log}
}

func (Like LikeDelivery) LikePhoto(rwContext echo.Context) error {
	photoId, err := strconv.Atoi(rwContext.Param("id"))
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.LikePhoto(photoId, userId)

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	Like.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}

func (Like LikeDelivery) DislikePhoto(rwContext echo.Context) error {
	photoId, err := strconv.Atoi(rwContext.Param("id"))
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.DislikePhoto(photoId, userId)

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	Like.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}

func (Like LikeDelivery) LikePost(rwContext echo.Context) error {
	postId, err := strconv.Atoi(rwContext.Param("id"))
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.LikePost(postId, userId)

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	Like.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}

func (Like LikeDelivery) DislikePost(rwContext echo.Context) error {
	postId, err := strconv.Atoi(rwContext.Param("id"))
	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.DislikePost(postId, userId)

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	Like.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}

func (Like LikeDelivery) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/photo/:id/like", Like.LikePhoto)
	server.DELETE("/api/v1/photo/:id/like", Like.DislikePhoto)

	server.POST("/api/v1/post/:id/like", Like.LikePost)
	server.DELETE("/api/v1/post/:id/like", Like.DislikePost)
}
