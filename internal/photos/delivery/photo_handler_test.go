package delivery

import (

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/internal/models"


	mock_photos "main/mocks"

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
	login := 1234

	usersId := []int{-1, 1 ,2}
	friendBehaviour := []error{nil, nil, errors.New("smth happend")}
	expectedBehaviour := []int{http.StatusUnauthorized, http.StatusNotFound, http.StatusNotFound}


				photoMock.EXPECT().UploadPhotoToAlbum(*photoData).Return(nil)
			}


		answer := models.Photos{}
		lUseCase.EXPECT().GetPhotosFromAlbum(login).Return(answer, friendBehaviour[iter])


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

		c.SetParamValues("1234")
		c.Set("REQUEST_ID", "123")
		c.Set("user_id", usersId[iter])


		if assert.NoError(t, photoTest.UploadPhotoToAlbum(c)) {
			assert.Equal(t, expectedBehaviour[iter], rec.Code)
		}
	}

}



