package delivery

import (
	"github.com/labstack/echo"
	"net/http/httptest"
	"strconv"
	"testing"
	"main/internal/microservices/photo_save/usecase"
)


func TestPhotoSaveDeliveryRealisation_Upload(t *testing.T) {
	photoSaverUseCase := usecase.NewUserUseCaseRealisation()
	photoSaverDeliveryTest := NewSavePhotoDeliveryRealisation(photoSaverUseCase)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("api/v1/albums/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))
	c.Set("REQUEST_ID", "123")
	c.Set("user_id", 1)

}
