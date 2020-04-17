package delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	mock_albums "main/mocks"
	"main/internal/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFriendDeliveryRealisation_GetAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_albums.NewMockAlbumUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	albumD := NewAlbumDelivery(logger, aUseCase)

	usersId := []int{-1, 1, 2}
	albumBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusInternalServerError}

	for iter, _ := range usersId {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			albums := make([]models.Album , 2 ,2 )
			aUseCase.EXPECT().GetAlbums(usersId[iter]).Return(albums, albumBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, albumD.GetAlbums(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}
}

func TestFriendDeliveryRealisation_CreateAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_albums.NewMockAlbumUseCase(ctrl)
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	albumD := NewAlbumDelivery(logger, aUseCase)

	usersId := []int{-1, 1, 2}
	albumBehaviour := []error{nil, nil, errors.New("smth happend") , nil , nil }
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusConflict , http.StatusInternalServerError}

	for iter, _ := range usersId {

		album := models.AlbumReq{Name: "123"}
		if expectedBehaviour[iter] != http.StatusUnauthorized {
			aUseCase.EXPECT().CreateAlbum(usersId[iter] , album).Return(albumBehaviour[iter])
		}

		e := echo.New()
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(`{"name" : "123"}`))
		req.Header.Set(echo.HeaderContentType , echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/album")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])

		if assert.NoError(t, albumD.CreateAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

	uId := 4

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	req.Header.Set(echo.HeaderContentType , echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/album")
	c.Set("REQUEST_ID", "123")
	c.Set("user_id", uId)

	if assert.NoError(t, albumD.CreateAlbum(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
}