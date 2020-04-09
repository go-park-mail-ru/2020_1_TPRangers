package albums

import (
	"main/internal/models"
)

type AlbumUseCase interface {
	GetAlbums(string) ([]models.Album, error)
	CreateAlbum(string, models.AlbumReq) error
}