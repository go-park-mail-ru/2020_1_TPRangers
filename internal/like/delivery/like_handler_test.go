package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	mock_like "main/mocks"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestLikeDelivery_LikePhoto(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		photoId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().LikePhoto(photoId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/photo/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(photoId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.LikePhoto(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}

func TestLikeDelivery_DislikePhoto(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		photoId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().DislikePhoto(photoId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/photo/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(photoId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.DislikePhoto(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}

func TestLikeDelivery_LikePost(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		postId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().LikePost(postId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/post/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(postId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.LikePost(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}

func TestLikeDelivery_DislikePost(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		postId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().DislikePost(postId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/post/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(postId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.DislikePost(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}


func TestLikeDelivery_LikeComment(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		postId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().LikeComment(postId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/post/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(postId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.LikeComment(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}

func TestLikeDelivery_DislikeComment(t *testing.T) {

	ctrl := gomock.NewController(t)
	lUseCase := mock_like.NewMockUseCaseLike(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	likeD := NewLikeDelivery(logger, lUseCase)

	usersId := []int{-1, 1, 2}
	likeBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict}

	for iter, _ := range usersId {
		postId := rand.Int()

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			lUseCase.EXPECT().DislikeComment(postId, usersId[iter]).Return(likeBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/post/:id/like")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(postId))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, likeD.DislikeComment(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

}