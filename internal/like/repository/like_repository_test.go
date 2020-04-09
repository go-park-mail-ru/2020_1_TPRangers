package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func recordLikePhotoStats(db *sql.DB, photoId int) (err error) {
	tx, err := db.Begin()

	if err != nil {
		fmt.Println("db not created" , err)
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _ , err = tx.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count + 1 WHERE photo_id = $1;" , photoId); err != nil {
		return
	}

	return
}

func TestLikeRepositoryRealisation_LikePhoto(t *testing.T) {

	db , mock, _ := sqlmock.New()
	testCounter := 4

	errs := []error{nil,sql.ErrNoRows, nil ,nil}


	for iter := 0  ; iter < testCounter ; iter ++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE Photos SET photos_likes_count \= photos_likes_count \+ 1 WHERE photo_id \= \$1; ` ).WithArgs(photoId).WillReturnResult(sqlmock.NewResult(0,int64(photoId)))

		mock.ExpectPrepare(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id;`)
		if errs[iter] == nil {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id;`).WithArgs(uId,photoId).WillReturnRows(sqlmock.NewRows([]string{"photolike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id;`).WithArgs(uId,photoId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()


		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created" , err)
			return
		}

		if _ , err = tx.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count + 1 WHERE photo_id = $1;" , photoId); err != nil {
			return
		}

		stmt , err := tx.Prepare("INSERT INTO UsersPhotosLikes (u_id,photo_id) VALUES ($1,$2) RETURNING photolike_id;")
		scanPLID := 0
		if  err = stmt.QueryRow(uId,photoId).Scan(&scanPLID); err != nil {
			return
		}

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}

}