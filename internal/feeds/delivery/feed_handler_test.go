package delivery

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/models"
	mock_feeds "main/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFriendDeliveryRealisation_Feed(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_feeds.NewMockFeedUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	feedD := NewFeedDelivery(logger, aUseCase)

	usersId := []int{-1, 1, 2}
	feedBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusInternalServerError}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			feeds := make([]models.Post, 2, 2)
			aUseCase.EXPECT().Feed(usersId[iter]).Return(feeds, feedBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, feedD.Feed(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_CreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	fUseCase := mock_feeds.NewMockFeedUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	feedD := NewFeedDelivery(logger, fUseCase)

	usersId := []int{-1, 1, 2}
	ownerLogin := "1234"
	feedBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusConflict, http.StatusConflict}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			post := models.Post{}
			fUseCase.EXPECT().CreatePost(usersId[iter], ownerLogin, post).Return(feedBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		fmt.Print(rec.Code)
		if assert.NoError(t, feedD.CreatePost(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_CreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	fUseCase := mock_feeds.NewMockFeedUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	feedD := NewFeedDelivery(logger, fUseCase)

	usersId := []int{-1, 1, 2}
	feedBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusConflict, http.StatusConflict}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			comment := models.Comment{}
			fUseCase.EXPECT().CreateComment(usersId[iter], comment).Return(feedBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, feedD.CreateComment(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

//func TestFriendDeliveryRealisation_GetPostAndComments(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	aUseCase := mock_feeds.NewMockFeedUseCase(ctrl)
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	feedD := NewFeedDelivery(logger, aUseCase)
//	postID := "1234"
//	usersId := []int{-1, 1, 2}
//	feedBehaviour := []error{nil, nil, errors.New("smth happend")}
//	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusInternalServerError}
//
//	for iter, _ := range usersId {
//
//		if expectedBehaviour[iter] != http.StatusUnauthorized {
//			post := models.Post{}
//			aUseCase.EXPECT().GetPostAndComments(usersId[iter], postID).Return(post, feedBehaviour[iter])
//		}
//
//		e := echo.New()
//		req := httptest.NewRequest(echo.GET, "/", nil)
//		rec := httptest.NewRecorder()
//		c := e.NewContext(req, rec)
//		c.Set("id", usersId[iter])
//		c.SetPath("/api/v1/post/:id/comments")
//		c.Set("REQUEST_ID", "123")
//		c.Set("user_id", usersId[iter])
//		if assert.NoError(t, feedD.GetPostAndComments(c)) {
//			assert.Equal(t, expectedBehaviour[iter], rec.Code)
//		}
//
//	}
//}