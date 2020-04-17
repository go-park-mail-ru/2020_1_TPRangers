package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"main/internal/models"
	"math/rand"
	"testing"
)

func TestLikeRepositoryRealisation_CreateAlbum(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 1

	lRepo := NewAlbumRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}

	for iter := 0; iter < testCounter; iter++ {

		uId := rand.Int()
		albumData := models.AlbumReq{}
		albumData.Name = "kek"
		mock.ExpectBegin()

		if errs[iter] == nil{
			mock.ExpectQuery(`INSERT INTO albums \(name, u_id\) VALUES \(\$1,\$2\)`).WithArgs(albumData.Name, uId)
		} else {
			mock.ExpectQuery(`INSERT INTO albums \(name, u_id\) VALUES \(\$1,\$2\)`).WithArgs(albumData.Name, uId).WillReturnError(errs[iter])
		}

		mock.ExpectCommit()

		tx, err := db.Begin()

		err = lRepo.CreateAlbum(uId, albumData)

		if err != errs[iter] {
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