package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"main/internal/models"
	"main/internal/tools/errors"
)

type AlbumRepositoryRealisation struct {
	albumDB *sql.DB
}

func NewAlbumRepositoryRealisation(db *sql.DB) AlbumRepositoryRealisation {
	return AlbumRepositoryRealisation{albumDB: db}

}

func (Data AlbumRepositoryRealisation) CreateAlbum(u_id int, albumData models.AlbumReq) error {

	_, err := Data.albumDB.Exec("INSERT INTO albums (name, u_id) VALUES ($1, $2);", albumData.Name, u_id)
	if err != nil {
		fmt.Print("ERR IS ", err)
		return errors.FailSendToDB
	}

	return nil

}

func (Data AlbumRepositoryRealisation) GetAlbums(id int) ([]models.Album, error) {
	albums := make([]models.Album, 0, 20)

	rows, err := Data.albumDB.Query("select DISTINCT ON (a.album_id) a.name, a.album_id, ph.photo_url from albums AS a LEFT JOIN photosfromalbums AS ph ON ph.album_id = a.album_id WHERE a.u_id = $1;", id)

	defer func() {
		if rows != nil {
			rows.Close()

		}
	}()

	if err != nil {
		return nil, errors.FailReadFromDB
	}

	for rows.Next() {
		var album models.Album
		err = rows.Scan(&album.Name, &album.ID, &album.PhotoUrl)

		if album.PhotoUrl == nil {
			album.PhotoUrl = new(string)
			*album.PhotoUrl = ""
		}
		if err != nil {
			return nil, errors.FailReadToVar
		}

		albums = append(albums, album)

	}
	return albums, nil
}
