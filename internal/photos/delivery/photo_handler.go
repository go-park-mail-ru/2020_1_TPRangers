package delivery

import (
	"fmt"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"main/internal/models"
	"main/internal/photos"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
)

type PhotoDeliveryRealisation struct {
	photoLogic photos.PhotoUseCase
	logger     *zap.SugaredLogger
}

func (photoD PhotoDeliveryRealisation) GetPhotosFromAlbum(rwContext echo.Context) error {

	uId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		photoD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	a_id, err := strconv.ParseInt(rwContext.Param("id"), 10, 32)

	photos, err := photoD.photoLogic.GetPhotosFromAlbum(int(a_id))

	if err != nil {
		photoD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	photoD.logger.Info(
		zap.String("ID", uId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	if len(photos.Urls) == 0 {
		return rwContext.JSON(http.StatusNotFound, models.JsonStruct{Body: photos})
	}

	return rwContext.JSON(http.StatusOK, photos)
}

func (photoD PhotoDeliveryRealisation) UploadPhotoToAlbum(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		photoD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	b , err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		photoD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	photoData := new(models.PhotoInAlbum)

	err = photoData.UnmarshalJSON(b)

	if err != nil {
		photoD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = photoD.photoLogic.UploadPhotoToAlbum(*photoData)

	fmt.Println(err)

	if err != nil {
		photoD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	photoD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}

func NewPhotoDelivery(log *zap.SugaredLogger, photoRealisation photos.PhotoUseCase) PhotoDeliveryRealisation {
	return PhotoDeliveryRealisation{photoLogic: photoRealisation, logger: log}
}

func (photoD PhotoDeliveryRealisation) InitHandlers(server *echo.Echo) {

	server.POST("api/v1/album/photo", photoD.UploadPhotoToAlbum)

	server.GET("api/v1/albums/:id", photoD.GetPhotosFromAlbum)

}
