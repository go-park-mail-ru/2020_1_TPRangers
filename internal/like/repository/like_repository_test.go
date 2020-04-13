package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
)

func TestLikeRepositoryRealisation_LikePhoto(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3

	lRepo := NewLikeRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}

	for iter := 0; iter < testCounter; iter++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnRows(sqlmock.NewRows([]string{"photolike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		db.Begin()
		err = lRepo.LikePhoto(photoId, uId)

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

func TestLikeRepositoryRealisation_DislikePhoto(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3
	lRepo := NewLikeRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}

	for iter := 0; iter < testCounter; iter++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`DELETE FROM UsersPhotosLikes WHERE u_id \= \$1 AND photo_id \= \$2 RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnRows(sqlmock.NewRows([]string{"photolike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`DELETE FROM UsersPhotosLikes WHERE u_id \= \$1 AND photo_id \= \$2 RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		db.Begin()
		err = lRepo.DislikePhoto(photoId, uId)

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

func TestLikeRepositoryRealisation_LikePost(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3
	lRepo := NewLikeRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}

	for iter := 0; iter < testCounter; iter++ {

		postId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(`INSERT INTO UsersPostsLikes \(u_id,post_id\) VALUES \(\$1,\$2\) RETURNING postlike_id`).WithArgs(uId, postId).WillReturnRows(sqlmock.NewRows([]string{"postlike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`INSERT INTO UsersPostsLikes \(u_id,post_id\) VALUES \(\$1,\$2\) RETURNING postlike_id`).WithArgs(uId, postId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		db.Begin()
		err = lRepo.LikePost(postId, uId)

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

func TestLikeRepositoryRealisation_DislikePost(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 3
	lRepo := NewLikeRepositoryRealisation(db)

	errs := []error{nil, sql.ErrNoRows, nil}

	for iter := 0; iter < testCounter; iter++ {

		postId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()

		if errs[iter] == nil {
			mock.ExpectQuery(` DELETE FROM UsersPostsLikes WHERE u_id \= \$1 AND post_id \= \$2 RETURNING postlike_id `).WithArgs(uId, postId).WillReturnRows(sqlmock.NewRows([]string{"postlike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(` DELETE FROM UsersPostsLikes WHERE u_id \= \$1 AND post_id \= \$2 RETURNING postlike_id `).WithArgs(uId, postId).WithArgs(uId, postId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		db.Begin()
		err = lRepo.DislikePost(postId, uId)

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
