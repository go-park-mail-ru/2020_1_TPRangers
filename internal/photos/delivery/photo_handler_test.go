package delivery

import (

	"crypto/rand"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/models"
	mock_photos "main/mocks"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFriendDeliveryRealisation_UploadPhotoToAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	aUseCase := mock_photos.NewMockPhotoUseCase(ctrl)

	"main/mocks"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestPhotoDeliveryRealisation_GetPhotosFromAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	photoMock := mock.NewMockPhotoUseCase(ctrl)

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

	photoTest := NewPhotoDelivery(logger, photoMock)

	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusNotFound, http.StatusConflict}
	users := []int{1, -1, 1, 1}

	for iter, _ := range expectedBehaviour {

		if expectedBehaviour[iter] != http.StatusUnauthorized {

			if expectedBehaviour[iter] == http.StatusOK {
				photoMock.EXPECT().GetPhotosFromAlbum(1).Return(models.Photos{
					AlbumName: "xd",
					Urls:      []string{"fuck"},
				}, nil)
			}

			if expectedBehaviour[iter] == http.StatusNotFound {

				photoMock.EXPECT().GetPhotosFromAlbum(1).Return(models.Photos{
					AlbumName: "xd",
					Urls:      nil,
				}, nil)

			}

			if expectedBehaviour[iter] == http.StatusConflict {
				photoMock.EXPECT().GetPhotosFromAlbum(1).Return(models.Photos{
					AlbumName: "xd",
					Urls:      []string{"fuck"},
				}, errors.New("123"))
			}

		}

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/albums/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(1))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t, photoTest.GetPhotosFromAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}

}

func TestPhotoDeliveryRealisation_UploadPhotoToAlbum(t *testing.T) {
	ctrl := gomock.NewController(t)
	photoMock := mock.NewMockPhotoUseCase(ctrl)

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	friendD := NewPhotoDelivery(logger, lUseCase)

	urls := []string{"first", "next", "third"}
	answer := models.Photos{
		AlbumName: "kek",
		Urls: urls,
	}


	usersId := []int{-1}
	friendBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusOK, http.StatusNotFound}

	for iter, _ := range usersId {

		login, _ := rand.Int(rand.Reader, big.NewInt(80))

		lUseCase.EXPECT().GetPhotosFromAlbum(login).Return(answer, friendBehaviour[iter])

		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/albums/:id")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])
		c.SetParamNames("id")
		c.SetParamValues(login.String())

		if assert.NoError(t, friendD.GetPhotosFromAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}

	}

	photoTest := NewPhotoDelivery(logger, photoMock)

	expectedBehaviour := []int{http.StatusOK, http.StatusUnauthorized, http.StatusConflict, http.StatusConflict, http.StatusConflict}
	users := []int{1, -1, 1, 1, 1}
	photos := []string{`{"url" : "private"}`, `{"url":"private"}`, `{"url":"private"}`, `haha`, `{"url":"private"}`}

	for iter, _ := range expectedBehaviour {

		if expectedBehaviour[iter] != http.StatusUnauthorized {
			if expectedBehaviour[iter] == http.StatusOK {
				photoData := new(models.PhotoInAlbum)
				photoData.Url = "private"

				photoMock.EXPECT().UploadPhotoToAlbum(*photoData).Return(nil)
			}

			if len(expectedBehaviour)-1 == iter {
				photoData := new(models.PhotoInAlbum)
				photoData.Url = "private"

				photoMock.EXPECT().UploadPhotoToAlbum(*photoData).Return(errors.New("123"))
			}
		}

		e := echo.New()
		var req *http.Request
		if iter == 2 {
			req = httptest.NewRequest(echo.GET, "/", nil)
		} else {
			req = httptest.NewRequest(echo.GET, "/", strings.NewReader(photos[iter]))
		}
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/albums/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(1))
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", users[iter])

		if assert.NoError(t, photoTest.UploadPhotoToAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}


}
