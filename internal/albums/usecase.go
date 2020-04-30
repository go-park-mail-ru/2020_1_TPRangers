package albums

import (
	"main/models"
)

type AlbumUseCase interface {
	GetAlbums(int) ([]models.Album, error)
	CreateAlbum(int, models.AlbumReq) error
}
