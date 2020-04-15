package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"main/internal/albums/usecase"
	"main/internal/models"
	"main/internal/albums"
	"main/internal/tools/errors"
	"net/http"
	"time"
)

type AlbumDeliveryRealisation struct {
	albumLogic albums.AlbumUseCase
	logger    *zap.SugaredLogger
}

func (albumD AlbumDeliveryRealisation) GetAlbums(rwContext echo.Context) error {
	rId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")
	if err != nil {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("COOKIE", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	albums, err := albumD.albumLogic.GetAlbums(cookie.Value)

	respErrStat := 0
	switch err {
	case errors.InvalidCookie:
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		rwContext.SetCookie(cookie)

		respErrStat = http.StatusUnauthorized
	case errors.FailReadFromDB:
		respErrStat = http.StatusInternalServerError
	}

	if err != nil {
		albumD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", respErrStat),
		)

		return rwContext.JSON(respErrStat, models.JsonStruct{Err: err.Error()})
	}

	albumD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.JSON(http.StatusOK, models.JsonStruct{Body: albums})
}

func (albumD AlbumDeliveryRealisation) CreateAlbum(rwContext echo.Context) error {
	rId := rwContext.Response().Header().Get("REQUEST_ID")

	cookie, err := rwContext.Cookie("session_id")

	if err != nil {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	albumData := new(models.AlbumReq)

	err = rwContext.Bind(albumData)

	if err != nil {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = albumD.albumLogic.CreateAlbum(cookie.Value, *albumData)

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
		albumD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", errRespStatus),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	albumD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}


func NewAlbumDelivery(log *zap.SugaredLogger, albumRealisation usecase.AlbumUseCaseRealisation) AlbumDeliveryRealisation {
	return AlbumDeliveryRealisation{albumLogic: albumRealisation, logger: log}
}

func (albumD AlbumDeliveryRealisation) InitHandlers(server *echo.Echo) {

	server.POST("api/v1/album", albumD.CreateAlbum)

	server.GET("api/v1/albums", albumD.GetAlbums)


}
