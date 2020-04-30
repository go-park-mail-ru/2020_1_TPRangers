package usecase

import (
	"main/internal/albums"
	"main/internal/models"
	"main/internal/tools/errors"
)

type AlbumUseCaseRealisation struct {
	albumDB albums.AlbumRepository
}

func (albumR AlbumUseCaseRealisation) GetAlbums(userId int) ([]models.Album, error) {

	albums, err := albumR.albumDB.GetAlbums(userId)

	return albums, err

}

func (albumR AlbumUseCaseRealisation) CreateAlbum(userId int, albumData models.AlbumReq) error {
	err := albumR.albumDB.CreateAlbum(userId, albumData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func NewAlbumUseCaseRealisation(userDB albums.AlbumRepository) AlbumUseCaseRealisation {
	return AlbumUseCaseRealisation{
		albumDB: userDB,
	}
}
