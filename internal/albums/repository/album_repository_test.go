package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"main/internal/models"
	"main/internal/tools/errors"
	"math/rand"
	"testing"
)

func TestAlbumRepositoryRealisation_CreateAlbum(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3

	lRepo := NewAlbumRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}
	expectBehavior := []error{nil, errors.FailSendToDB, nil}
	for iter := 0; iter < testCounter; iter++ {

		uId := rand.Int()
		albumData := models.AlbumReq{}
		albumData.Name = "kek"
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectExec(`INSERT INTO albums \(name, u_id\) VALUES \(\$1, \$2\);`).WithArgs(albumData.Name, uId).WillReturnResult(sqlmock.NewResult(1, 1))
		} else {
			mock.ExpectExec(`INSERT INTO albums \(name, u_id\) VALUES \(\$1, \$2\);`).WithArgs(albumData.Name, uId).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()

		err = lRepo.CreateAlbum(uId, albumData)

		if err != expectBehavior[iter] {
			fmt.Print(err)
			t.Error(err)
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

func TestAlbumRepositoryRealisation_GetAlbums(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3

	lRepo := NewAlbumRepositoryRealisation(db)

	errs := []error{nil, errors.FailSendToDB, nil}
	expectBehavior := []error{nil, errors.FailReadFromDB, errors.FailReadToVar}
	for iter := 0; iter < testCounter; iter++ {

		Id := rand.Int()

		mock.ExpectBegin()
		albStr := models.Album{}
		if errs[iter] == nil {
			if expectBehavior[iter] != errors.FailReadToVar {
				mock.ExpectQuery(`select DISTINCT ON \(a\.album_id\) a\.name, a\.album_id, ph\.photo_url from albums AS a LEFT JOIN photosfromalbums AS ph ON ph\.album_id \= a\.album_id WHERE a\.u_id \= \$1;`).WithArgs(Id).WillReturnRows(sqlmock.NewRows([]string{"a.name", "a.album_id", "ph.photo_url"}).AddRow(albStr.Name, albStr.ID, albStr.PhotoUrl))
			} else {
				mock.ExpectQuery(`select DISTINCT ON \(a\.album_id\) a\.name, a\.album_id, ph\.photo_url from albums AS a LEFT JOIN photosfromalbums AS ph ON ph\.album_id \= a\.album_id WHERE a\.u_id \= \$1;`).WithArgs(Id).WillReturnRows(sqlmock.NewRows([]string{"a.name", "a.album_id", "ph.photo_url"}).AddRow(nil, albStr.ID, albStr.PhotoUrl))
			}
		} else {
			mock.ExpectQuery(`select DISTINCT ON \(a\.album_id\) a\.name, a\.album_id, ph\.photo_url from albums AS a LEFT JOIN photosfromalbums AS ph ON ph\.album_id \= a\.album_id WHERE a\.u_id \= \$1;`).WithArgs(Id).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()

		_, err = lRepo.GetAlbums(Id)

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
