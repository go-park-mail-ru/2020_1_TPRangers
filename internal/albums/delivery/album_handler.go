package delivery

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"main/internal/albums"
	"main/internal/models"
	"main/internal/tools/errors"
	"net/http"
)

type AlbumDeliveryRealisation struct {
	albumLogic albums.AlbumUseCase
	logger     *zap.SugaredLogger
}

func (albumD AlbumDeliveryRealisation) GetAlbums(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("COOKIE", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)

		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	albums, err := albumD.albumLogic.GetAlbums(userId)

	if err != nil {
		albumD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.JSON(http.StatusInternalServerError, models.JsonStruct{Err: err.Error()})
	}

	albumD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)
	return rwContext.JSON(http.StatusOK, albums)
}

func (albumD AlbumDeliveryRealisation) CreateAlbum(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}
	albumData := new(models.AlbumReq)

	b, err := ioutil.ReadAll(rwContext.Request().Body)
	defer rwContext.Request().Body.Close()

	if err != nil {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = albumData.UnmarshalJSON(b)

	if err != nil {
		albumD.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)

		return rwContext.NoContent(http.StatusInternalServerError)
	}

	err = albumD.albumLogic.CreateAlbum(userId, *albumData)

	if err != nil {
		albumD.logger.Info(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	albumD.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.NoContent(http.StatusOK)
}

func NewAlbumDelivery(log *zap.SugaredLogger, albumRealisation albums.AlbumUseCase) AlbumDeliveryRealisation {
	return AlbumDeliveryRealisation{albumLogic: albumRealisation, logger: log}
}

func (albumD AlbumDeliveryRealisation) InitHandlers(server *echo.Echo) {

	server.POST("api/v1/album", albumD.CreateAlbum)

	server.GET("api/v1/albums", albumD.GetAlbums)

}
