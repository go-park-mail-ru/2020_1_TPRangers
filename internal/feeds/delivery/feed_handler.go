package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/feeds"
	"main/internal/feeds/usecase"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
)

type FeedDeliveryRealisation struct {
	feedLogic feeds.FeedUseCase
	logger    *zap.SugaredLogger
}

func (feedD FeedDeliveryRealisation) Feed(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := feedD.feedLogic.Feed(userId)

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func (feedD FeedDeliveryRealisation) CreatePost(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)
	ownerLogin := rwContext.Param("id")

	if userId == -1 {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	newPost := new(models.Post)

	err := rwContext.Bind(&newPost)

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}
	err = feedD.feedLogic.CreatePost(userId, ownerLogin, *newPost)
	if err != nil {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func (feedD FeedDeliveryRealisation) CreateComment(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	newComment := new(models.Comment)
	err := rwContext.Bind(&newComment)

	if err != nil {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)
		return rwContext.JSON(http.StatusConflict, models.JsonStruct{Err: err.Error()})
	}

	err = feedD.feedLogic.CreateComment(userId,  *newComment)

	if err != nil {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}
	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.NoContent(http.StatusOK)
}
func (feedD FeedDeliveryRealisation) DeleteComment(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	commentID := rwContext.Param("id")
	if userId == -1 {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	err := feedD.feedLogic.DeleteComment(userId, commentID)
	if err == errors.DontHavePermission {
		return rwContext.JSON(http.StatusForbidden, models.JsonStruct{Err: err.Error()})
	}
	if err != nil {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}
	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.NoContent(http.StatusOK)
}

func (feedD FeedDeliveryRealisation) GetComments(rwContext echo.Context) error {
	uId := rwContext.Get("REQUEST_ID").(string)
	userId := rwContext.Get("user_id").(int)
	postID := rwContext.Param("id")
	if userId == -1 {
		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	jsonAnswer, err := feedD.feedLogic.GetComments(userId, postID)

	if err != nil {

		feedD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	feedD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, jsonAnswer)
}

func NewFeedDelivery(log *zap.SugaredLogger, feedRealisation usecase.FeedUseCaseRealisation) FeedDeliveryRealisation {
	return FeedDeliveryRealisation{feedLogic: feedRealisation, logger: log}
}

func (feedD FeedDeliveryRealisation) InitHandlers(server *echo.Echo) {
	server.GET("/api/v1/news", feedD.Feed)
	server.POST("/api/v1/:id/post", feedD.CreatePost)
	server.POST("/api/v1/comment", feedD.CreateComment)
	server.DELETE("/api/v1/comment/:id", feedD.DeleteComment)
	server.GET("/api/v1/post/:id/comments", feedD.GetComments)
}
