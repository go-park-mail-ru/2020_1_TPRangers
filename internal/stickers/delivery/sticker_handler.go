package delivery

import (
	"main/internal/models"
	"main/internal/stickers"
	"main/internal/tools/errors"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

type StickerDeliveryRealisation struct {
	stickerLogic stickers.StickerUse
	logger       *zap.SugaredLogger
}

func NewStickerDelivery(log *zap.SugaredLogger, stickerRealisation stickers.StickerUse) StickerDeliveryRealisation {
	return StickerDeliveryRealisation{stickerLogic: stickerRealisation, logger: log}
}

func (Sticker StickerDeliveryRealisation) UploadPack(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	pack := new(models.StickerPack)
	err := rwContext.Bind(&pack)

	if err != nil {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = Sticker.stickerLogic.CreateNewPack(userId, *pack)

	if err != nil {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	Sticker.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusCreated),
	)

	return rwContext.NoContent(http.StatusCreated)
}

func (Sticker StickerDeliveryRealisation) PurchasePack(rwContext echo.Context) error {

	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	pack := new(models.StickerPack)
	err := rwContext.Bind(&pack)

	if err != nil {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	err = Sticker.stickerLogic.PurchasePack(userId, *pack)

	if err != nil {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	Sticker.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusCreated),
	)

	return rwContext.NoContent(http.StatusCreated)

}

func (Sticker StickerDeliveryRealisation) GetStickerPacks(rwContext echo.Context) error {
	rId := rwContext.Get("REQUEST_ID").(string)

	userId := rwContext.Get("user_id").(int)

	if userId == -1 {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", errors.CookieExpired.Error()),
			zap.Int("ANSWER STATUS", http.StatusUnauthorized),
		)
		return rwContext.JSON(http.StatusUnauthorized, models.JsonStruct{Err: errors.CookieExpired.Error()})
	}

	packs, err := Sticker.stickerLogic.GetStickerPacks(userId)

	if err != nil {
		Sticker.logger.Debug(
			zap.String("ID", rId),
			zap.String("ERROR", err.Error()),
			zap.Int("ANSWER STATUS", http.StatusConflict),
		)

		return rwContext.NoContent(http.StatusConflict)
	}

	Sticker.logger.Info(
		zap.String("ID", rId),
		zap.Int("ANSWER STATUS", http.StatusOK),
	)

	return rwContext.JSON(http.StatusOK, packs)
}

func (Sticker StickerDeliveryRealisation) InitHandlers(server *echo.Echo) {

	server.POST("api/v1/stickers", Sticker.UploadPack)
	server.PUT("api/v1/stickers", Sticker.PurchasePack)
	server.GET("api/v1/stickers", Sticker.GetStickerPacks)

}
