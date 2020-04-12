package usecase

import (
	"main/internal/albums"
	AlbumRep "main/internal/albums/repository"
	"main/internal/models"
	"main/internal/tools/errors"
)

type AlbumUseCaseRealisation struct {
	albumDB albums.AlbumRepository
}

func (albumR AlbumUseCaseRealisation) GetAlbums(userId int) ([]models.Album, error) {

	albums, _ := albumR.albumDB.GetAlbums(userId)

	return albums, nil

}

func (albumR AlbumUseCaseRealisation) CreateAlbum(userId int, albumData models.AlbumReq) error {
	err := albumR.albumDB.CreateAlbum(userId, albumData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}

func NewAlbumUseCaseRealisation(userDB AlbumRep.AlbumRepositoryRealisation) AlbumUseCaseRealisation {
	return AlbumUseCaseRealisation{
		albumDB: userDB,
	}
}
