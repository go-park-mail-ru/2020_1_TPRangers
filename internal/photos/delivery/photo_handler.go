package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/models"
	"main/internal/photos/usecase"
	"main/internal/photos"
	"main/internal/tools/errors"
	"net/http"
	"strconv"
	"time"
)


type PhotoDeliveryRealisation struct {
	photoLogic photos.PhotoUseCase
	logger    *zap.SugaredLogger
}

func (photoD PhotoDeliveryRealisation) GetPhotosFromAlbum(rwContext echo.Context) error {

	uId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		photoD.logger.Debug(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	a_id, err := strconv.ParseInt(rwContext.Param("id"), 10, 32)

	photos, err := photoD.photoLogic.GetPhotosFromAlbum(cookie.Value, int(a_id))

	errRespStatus := 0

	switch err {
	case errors.InvalidCookie:
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		errRespStatus = http.StatusUnauthorized
	case errors.FailReadFromDB:
		errRespStatus = http.StatusInternalServerError
	}

	if err != nil {
		photoD.logger.Info(
			zap.String("ID", uId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStatus),
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

	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: photos})
}



func (photoD PhotoDeliveryRealisation) UploadPhotoToAlbum(rwContext echo.Context) error {
	rId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		photoD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	photoData := new(models.PhotoInAlbum)

	err = rwContext.Bind(photoData)

	if err != nil {
		photoD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = photoD.photoLogic.UploadPhotoToAlbum(cookie.Value, *photoData)

	errRespStatus := 0

	switch err {
	case errors.InvalidCookie:
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)
		errRespStatus = http.StatusUnauthorized
	case errors.FailReadFromDB:
		errRespStatus = http.StatusInternalServerError
	}

	if err != nil {
		photoD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStatus),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	photoD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)

}





func NewPhotoDelivery(log *zap.SugaredLogger, photoRealisation usecase.PhotoUseCaseRealisation) PhotoDeliveryRealisation {
	return PhotoDeliveryRealisation{photoLogic: photoRealisation, logger: log}
}

func (photoD PhotoDeliveryRealisation) InitHandlers(server *echo.Echo) {

	server.POST("api/v1/album/photo", photoD.UploadPhotoToAlbum)


	server.GET("api/v1/albums/:id", photoD.GetPhotosFromAlbum)

}