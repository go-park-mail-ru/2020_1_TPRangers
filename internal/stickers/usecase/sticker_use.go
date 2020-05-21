package usecase

import (
	"main/internal/models"
	"main/internal/stickers"
)

type StickerUseRealisation struct {
	stickerRepo stickers.StickerRepo
}

func NewStickerUseRealisation(repo stickers.StickerRepo) StickerUseRealisation {
	return StickerUseRealisation{stickerRepo: repo}
}

func (Sticker StickerUseRealisation) CreateNewPack(userId int, pack models.StickerPack) error {
	return Sticker.stickerRepo.UploadStickerPack(userId, pack)
}

func (Sticker StickerUseRealisation) GetStickerPacks(userId int) ([]models.StickerPack, error) {
	return Sticker.stickerRepo.GetStickerPacks(userId)
}

func (Sticker StickerUseRealisation) PurchasePack(userId int, pack models.StickerPack) error {
	return Sticker.stickerRepo.PurchaseStickerPack(userId, *pack.PackId)
}
