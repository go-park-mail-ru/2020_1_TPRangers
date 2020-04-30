package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/models"
	mock_photos "main/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFriendDeliveryRealisation_UploadPhotoToAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_photos.NewMockPhotoUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	albumD := NewPhotoDelivery(logger, aUseCase)

	usersId := []int{-1, 1, 2}
	albumBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusInternalServerError, http.StatusInternalServerError}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			photo := models.PhotoInAlbum{}
			aUseCase.EXPECT().UploadPhotoToAlbum(photo).Return(albumBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album/photo")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, albumD.UploadPhotoToAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestPhotosDeliveryRealisation_GetPhotosFromAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	lUseCase := mock_photos.NewMockPhotoUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	friendD := NewPhotoDelivery(logger, lUseCase)
	login := 1234

	usersId := []int{-1, 1 ,2}
	friendBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusNotFound, http.StatusNotFound}

	for iter, _ := range usersId {

		answer := models.Photos{}
		lUseCase.EXPECT().GetPhotosFromAlbum(login).Return(answer, friendBehaviour[iter])

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/albums/:id")
		c.SetParamNames("id")
		c.SetParamValues("1234")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, friendD.GetPhotosFromAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}


