package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"main/internal/models"
	"main/internal/tools/errors"
	"testing"
)


func TestPhotoRepositoryRealisation_UploadPhotoToAlbum(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 1

	lRepo := NewPhotoRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows}
	expectBehavior := []error{errors.AlbumDoesntExist, nil}
	for iter := 0; iter < testCounter; iter++ {

		photoInAlb:= models.PhotoInAlbum{}
		photoInAlb.AlbumID = "1"
		photoInAlb.Url = "kek"
		mock.ExpectBegin()
		if errs[iter] == nil{
			mock.ExpectExec(`select name from albums where album_id \= \$1;`).WithArgs(photoInAlb.AlbumID).WillReturnResult(sqlmock.NewResult(1,1))
		} else {
			mock.ExpectExec(`select name from albums where album_id \= \$1;`).WithArgs(photoInAlb.AlbumID).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()
		tx, err := db.Begin()
		err = lRepo.UploadPhotoToAlbum(photoInAlb)

		if err != expectBehavior[iter] {
			fmt.Print(err)
			t.Error(err)
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
	testCounter := 1

	Id := 0
	ph_url := new(string)
	errs := []error{nil, sql.ErrNoRows}
	expectBehavior := []error{nil, nil}
	for iter := 0; iter < testCounter; iter++ {

		mock.ExpectBegin()
		if errs[iter] == nil{
			mock.ExpectQuery(`select photo_url from photosfromalbums where album_id \= \$1;`).WithArgs(Id).WillReturnRows(sqlmock.NewRows([]string{"photo_url"}).AddRow(ph_url))
		} else {
			mock.ExpectQuery(`select photo_url from photosfromalbums where album_id \= \$1;`).WithArgs(Id).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()
		tx, err := db.Begin()

		if err != expectBehavior[iter] {
			fmt.Print(err)
			t.Error(err)
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

