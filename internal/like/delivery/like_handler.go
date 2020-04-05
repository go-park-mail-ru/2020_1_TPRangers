package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/like"
	"main/internal/like/usecase"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
	"time"
)

type LikeDelivery struct {
	likeLogic like.UseCaseLike
	logger    *zap.SugaredLogger
}

func NewLikeDelivery(log *zap.SugaredLogger, likeRealisation usecase.LikesUseRealisation) LikeDelivery {
	return LikeDelivery{likeLogic: likeRealisation, logger: log}
}

func (Like LikeDelivery) LikePhoto(rwContext echo.Context) error {
	photoId, err := strconv.Atoi(rwContext.Param("id"))
	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie , err := rwContext.Cookie("session_id")

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.LikePhoto(photoId, cookie.Value)

	if err == errors.InvalidCookie{
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: err.Error()})
	}

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
	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie , err := rwContext.Cookie("session_id")

	if err != nil {
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err = Like.likeLogic.DislikePhoto(photoId, cookie.Value)

	if err == errors.InvalidCookie{
		Like.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: err.Error()})
	}

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


func (likeD LikeDelivery) InitHandlers(server *echo.Echo) {
	server.POST("/api/v1/photo/:id/like", likeD.LikePhoto)
	server.DELETE("/api/v1/photo/:id/like", likeD.DislikePhoto)


}
