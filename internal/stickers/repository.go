package stickers

import "main/internal/models"

type StickerRepo interface {
	UploadStickerPack(int, models.StickerPack) error
	GetStickerPacks(int) ([]models.StickerPack, error)
	PurchaseStickerPack(int, int64) error
}
