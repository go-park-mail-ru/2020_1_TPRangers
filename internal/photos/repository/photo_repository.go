package repository

import (
	"database/sql"
	"fmt"
	"main/internal/models"
	"main/internal/tools/errors"
	"strconv"
)

type PhotoRepositoryRealisation struct {
	photoDB *sql.DB
}

func NewPhotoRepositoryRealisation(db *sql.DB) PhotoRepositoryRealisation {
	return PhotoRepositoryRealisation{photoDB: db}

}

func (Data PhotoRepositoryRealisation) UploadPhotoToAlbum(photoData models.PhotoInAlbum) error {
	albumId, err := strconv.ParseInt(photoData.AlbumID, 10, 32)

	album := Data.photoDB.QueryRow("select name from albums where album_id = $1;", int(albumId))
	var albumName string
	err = album.Scan(&albumName)
	if albumName == "" {
		return errors.AlbumDoesntExist
	}

	_, err = Data.photoDB.Exec("INSERT INTO photos (url, photos_likes_count) VALUES ($1, $2);", photoData.Url, 0)
	if err != nil {
		//fmt.Println(err)
		return errors.FailSendToDB
	}
	var photoID int
	row := Data.photoDB.QueryRow("select photo_id from photos where url = $1", photoData.Url)
	err = row.Scan(&photoID)
	if err != nil {
		return errors.FailReadToVar
	}

	_, err = Data.photoDB.Exec("INSERT INTO photosfromalbums (photo_id, photo_url, album_id) VALUES ($1, $2, $3);", photoID, photoData.Url, int(albumId))
	if err != nil {
		return errors.FailSendToDB
	}
	return nil
}

func (Data PhotoRepositoryRealisation) GetPhotosFromAlbum(albumID int) (models.Photos, error) {
	photosAlb := models.Photos{}
	phUrls := make([]string, 0, 20)
	rows, err := Data.photoDB.Query("select photo_url from photosfromalbums where album_id = $1;", albumID)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	} ()
	if err != nil {
		fmt.Println(err)
		return models.Photos{}, errors.FailReadFromDB
	}

	for rows.Next() {
		var phUrl string

		err = rows.Scan(&phUrl)

		if err != nil {
			return models.Photos{}, errors.FailReadToVar
		}

		phUrls = append(phUrls, phUrl)
	}
	photosAlb.Urls = phUrls
	row, err := Data.photoDB.Query("select name from albums where album_id = $1;", albumID)
	err = row.Scan(&photosAlb.AlbumName)
	if err != nil {
		return models.Photos{}, nil
	}

	return photosAlb, nil
}
