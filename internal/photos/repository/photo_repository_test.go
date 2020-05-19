package repository

import (
	"database/sql"
	errors2 "errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"main/internal/models"
	"main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestPhotoRepositoryRealisation_UploadPhotoToAlbum(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 5

	lRepo := NewPhotoRepositoryRealisation(db)

	customErr := errors2.New("smth wrong")
	errs := []error{errors.AlbumDoesntExist, errors.FailSendToDB, errors.FailReadToVar, customErr, nil}
	expectBehavior := []error{errors.AlbumDoesntExist, errors.FailSendToDB, errors.FailReadToVar, errors.FailSendToDB, nil}
	for iter := 0; iter < testCounter; iter++ {

		photoInAlb := models.PhotoInAlbum{}
		photoInAlb.AlbumID = "1"
		photoInAlb.Url = "kek"
		photoId := rand.Int()
		mock.ExpectBegin()
		if expectBehavior[iter] != errors.AlbumDoesntExist {
			mock.ExpectQuery(`select name from albums where album_id \= \$1;`).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("new name"))

			if expectBehavior[iter] != errors.FailSendToDB {
				mock.ExpectExec(` INSERT INTO photos \(url, photos_likes_count\) VALUES \(\$1, \$2\); `).WithArgs(photoInAlb.Url, 0).WillReturnResult(sqlmock.NewResult(1, 1))

				if expectBehavior[iter] != errors.FailReadToVar {
					mock.ExpectQuery(` select photo_id from photos where url \= \$1 `).WithArgs(photoInAlb.Url).WillReturnRows(sqlmock.NewRows([]string{"photo_id"}).AddRow(photoId))

					if errs[iter] != customErr {
						mock.ExpectExec(` INSERT INTO photosfromalbums \(photo_id, photo_url, album_id\) VALUES \(\$1, \$2, \$3\); `).WithArgs(photoId, photoInAlb.Url, 1).WillReturnResult(sqlmock.NewResult(1, 1))
					} else {
						mock.ExpectExec(` INSERT INTO photosfromalbums \(photo_id, photo_url, album_id\) VALUES \(\$1, \$2, \$3\); `).WithArgs(photoId, photoInAlb.Url, 1).WillReturnError(customErr)
					}

				} else {
					mock.ExpectQuery(` select photo_id from photos where url \= \$1 `).WithArgs(photoInAlb.Url).WillReturnError(errs[iter])
				}

			} else {
				mock.ExpectExec(` INSERT INTO photos \(url, photos_likes_count\) VALUES \(\$1, \$2\); `).WithArgs(photoInAlb.Url, 0).WillReturnError(errs[iter])
			}

		} else {
			mock.ExpectQuery(`select name from albums where album_id \= \$1;`).WithArgs(1).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()
		tx, err := db.Begin()
		if err != nil {
			return
		}
		err = lRepo.UploadPhotoToAlbum(photoInAlb)

		if err != expectBehavior[iter] {
			fmt.Print(iter, err, expectBehavior[iter])
			return
		}
		err = nil
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}
}

func TestPhotoRepositoryRealisation_GetPhotosFromAlbum(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 2

	ph_url := new(string)
	errs := []error{nil, sql.ErrNoRows}
	expectBehavior := []error{nil, errors.FailReadFromDB}
	lRepo := NewPhotoRepositoryRealisation(db)

	for iter := 0; iter < testCounter; iter++ {

		Id := rand.Int()

		mock.ExpectBegin()
		if errs[iter] == nil {
			mock.ExpectQuery(`select photo_url from photosfromalbums where album_id \= \$1;`).WithArgs(Id).WillReturnRows(sqlmock.NewRows([]string{"photo_url"}).AddRow(ph_url))

			mock.ExpectQuery(` select name from albums where album_id \= \$1 `).WithArgs(Id).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("213"))

		} else {
			mock.ExpectQuery(`select photo_url from photosfromalbums where album_id \= \$1;`).WithArgs(Id).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()
		tx, err := db.Begin()

		if _, err = lRepo.GetPhotosFromAlbum(Id); err != expectBehavior[iter] {
			t.Error(iter, err, expectBehavior[iter])
			return
		}
		err = nil
		switch err {
		case nil:
			err = tx.Commit()
			if err != nil {
				return
			}
		default:
			err = tx.Rollback()
			if err != nil {
				return
			}
		}
	}
}
