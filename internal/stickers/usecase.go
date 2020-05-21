package stickers

import "main/internal/models"

type StickerUse interface {
	CreateNewPack(int, models.StickerPack) error
	GetStickerPacks(int) ([]models.StickerPack, error)
	PurchasePack(int, models.StickerPack) error
}
