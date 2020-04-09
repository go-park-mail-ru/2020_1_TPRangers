package usecase

import (
	"fmt"
	Sess "main/internal/cookies"
	SessRep "main/internal/cookies/repository"
	"main/internal/models"
	"main/internal/tools/errors"
	"main/internal/albums"
	AlbumRep "main/internal/albums/repository"

)

type AlbumUseCaseRealisation struct {
	albumDB   albums.AlbumRepository
	sessionDB Sess.CookieRepository
}

func (albumR AlbumUseCaseRealisation) GetAlbums(cookie string) ([]models.Album, error) {
	id, err := albumR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return  nil ,errors.InvalidCookie
	}

	albums, err := albumR.albumDB.GetAlbums(id)

	fmt.Println(albums)

	return albums , nil

}



func (albumR AlbumUseCaseRealisation) CreateAlbum(cookie string, albumData models.AlbumReq) error {
	uID, err := albumR.sessionDB.GetUserIdByCookie(cookie)

	if err != nil {
		return errors.InvalidCookie
	}

	err = albumR.albumDB.CreateAlbum(uID, albumData)

	if err != nil {
		return errors.FailReadFromDB
	}

	return nil
}



func NewAlbumUseCaseRealisation(userDB AlbumRep.AlbumRepositoryRealisation, sesDB SessRep.CookieRepositoryRealisation) AlbumUseCaseRealisation {
	return AlbumUseCaseRealisation{
		albumDB:    userDB,
		sessionDB: sesDB,
	}
}

