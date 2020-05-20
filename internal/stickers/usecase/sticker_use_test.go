package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"main/internal/models"
	mock "main/mocks"
	"testing"
)


func TestStickerUseRealisation_CreateNewPack(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockStickerRepo(ctrl)
	stickerUse := NewStickerUseRealisation(repoMock)
	customErr := errors.New("smth happend")
	pack := models.StickerPack{
		PackId:   nil,
		Author:   nil,
		Name:     nil,
		Readme:   nil,
		Stickers: []models.Sticker{},
		Owned:    false,
	}
	userId := 20
	repoMock.EXPECT().UploadStickerPack(userId,pack).Return(customErr)

	if err := stickerUse.CreateNewPack(userId,pack); err != customErr {
		t.Error("wrong behaviour")
	}
}

func TestStickerUseRealisation_GetStickerPacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockStickerRepo(ctrl)
	stickerUse := NewStickerUseRealisation(repoMock)
	customErr := errors.New("smth happend")
	packs := []models.StickerPack{ models.StickerPack{
		PackId:   nil,
		Author:   nil,
		Name:     nil,
		Readme:   nil,
		Stickers: []models.Sticker{},
		Owned:    false,
	},

	}
	userId := 20
	repoMock.EXPECT().GetStickerPacks(userId).Return(packs,customErr)

	if _ ,err := stickerUse.GetStickerPacks(userId); err != customErr {
		t.Error("wrong behaviour")
	}
}

func TestStickerUseRealisation_PurchasePack(t *testing.T) {
	ctrl := gomock.NewController(t)
	repoMock := mock.NewMockStickerRepo(ctrl)
	stickerUse := NewStickerUseRealisation(repoMock)
	customErr := errors.New("smth happend")
	packId := int64(40)
	pack := models.StickerPack{
		PackId:   &packId,
		Author:   nil,
		Name:     nil,
		Readme:   nil,
		Stickers: []models.Sticker{},
		Owned:    false,
	}
	userId := 20
	repoMock.EXPECT().PurchaseStickerPack(userId,packId).Return(customErr)

	if err := stickerUse.PurchasePack(userId,pack); err != customErr {
		t.Error("wrong behaviour")
	}
}
