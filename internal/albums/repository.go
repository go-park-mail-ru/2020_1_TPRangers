package albums

import "main/internal/models"

type AlbumRepository interface {
	GetAlbums(int) ([]models.Album, error)
	CreateAlbum(int, models.AlbumReq) error
}
