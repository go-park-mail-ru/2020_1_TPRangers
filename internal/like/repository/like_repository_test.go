package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"math/rand"
	"testing"
)

func TestLikeRepositoryRealisation_LikePhoto(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 4

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE Photos SET photos_likes_count \= photos_likes_count \+ 1 WHERE photo_id \= \$1 `).WithArgs(photoId).WillReturnResult(sqlmock.NewResult(0, int64(photoId)))

		mock.ExpectPrepare(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id`)
		if errs[iter] == nil {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnRows(sqlmock.NewRows([]string{"photolike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`INSERT INTO UsersPhotosLikes \(u_id,photo_id\) VALUES \(\$1,\$2\) RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		if _, err = tx.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count + 1 WHERE photo_id = $1", photoId); err != nil {
			return
		}

		stmt, err := tx.Prepare("INSERT INTO UsersPhotosLikes (u_id,photo_id) VALUES ($1,$2) RETURNING photolike_id")
		scanPLID := 0
		if err = stmt.QueryRow(uId, photoId).Scan(&scanPLID); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}
		}

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
	testCounter := 4

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE Photos SET photos_likes_count \= photos_likes_count \+ 1 WHERE photo_id \= \$1`).WithArgs(photoId).WillReturnResult(sqlmock.NewResult(0, int64(photoId)))

		if errs[iter] == nil {
			mock.ExpectExec(`DELETE FROM UsersPhotosLikes WHERE u_id \= \$1 AND photo_id \= \$2 RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnResult(sqlmock.NewResult(0, int64(likeId)))
		} else {
			mock.ExpectExec(`DELETE FROM UsersPhotosLikes WHERE u_id \= \$1 AND photo_id \= \$2 RETURNING photolike_id`).WithArgs(uId, photoId).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		if _, err = tx.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count + 1 WHERE photo_id = $1", photoId); err != nil {
			tx.Rollback()
		}

		if _, err = tx.Exec("DELETE FROM UsersPhotosLikes WHERE u_id = $1 AND photo_id = $2 RETURNING photolike_id", uId, photoId); err != nil {
			tx.Rollback()
		}

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
	testCounter := 4

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {

		postId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()
		mock.ExpectExec(` UPDATE Posts SET posts_likes_count \= posts_likes_count \+ 1 WHERE post_id \=\$1 `).WithArgs(postId).WillReturnResult(sqlmock.NewResult(0, int64(postId)))

		mock.ExpectPrepare(`INSERT INTO UsersPostsLikes \(u_id,post_id\) VALUES \(\$1,\$2\) RETURNING postlike_id`)
		if errs[iter] == nil {
			mock.ExpectQuery(`INSERT INTO UsersPostsLikes \(u_id,post_id\) VALUES \(\$1,\$2\) RETURNING postlike_id`).WithArgs(uId, postId).WillReturnRows(sqlmock.NewRows([]string{"photolike_id"}).AddRow(likeId))
		} else {
			mock.ExpectQuery(`INSERT INTO UsersPostsLikes \(u_id,post_id\) VALUES \(\$1,\$2\) RETURNING postlike_id`).WithArgs(uId, postId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		if _, err = tx.Exec("UPDATE Posts SET posts_likes_count = posts_likes_count + 1 WHERE post_id =$1", postId); err != nil {
			return
		}

		stmt, err := tx.Prepare("INSERT INTO UsersPostsLikes (u_id,post_id) VALUES ($1,$2) RETURNING postlike_id")
		scanPLID := 0
		if err = stmt.QueryRow(uId, postId).Scan(&scanPLID); err != nil {
			if err == sql.ErrNoRows {
				err = nil
			}

		}

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
	testCounter := 4

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {

		photoId := rand.Int()
		uId := rand.Int()
		likeId := rand.Int()
		mock.ExpectBegin()
		mock.ExpectExec(` UPDATE Posts SET posts_likes_count \= posts_likes_count \- 1 WHERE post_id \= \$1 `).WithArgs(photoId).WillReturnResult(sqlmock.NewResult(0, int64(photoId)))

		if errs[iter] == nil {
			mock.ExpectExec(` DELETE FROM UsersPostsLikes WHERE u_id \= \$1 AND post_id \= \$2 RETURNING postlike_id `).WithArgs(uId, photoId).WillReturnResult(sqlmock.NewResult(0, int64(likeId)))
		} else {
			mock.ExpectExec(` DELETE FROM UsersPostsLikes WHERE u_id \= \$1 AND post_id \= \$2 RETURNING postlike_id `).WithArgs(uId, photoId).WillReturnResult(sqlmock.NewResult(0, 0))
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		if _, err = tx.Exec("UPDATE Posts SET posts_likes_count = posts_likes_count - 1 WHERE post_id = $1", photoId); err != nil {
			tx.Rollback()
		}

		if _, err = tx.Exec("DELETE FROM UsersPostsLikes WHERE u_id = $1 AND post_id = $2 RETURNING postlike_id", uId, photoId); err != nil {
			tx.Rollback()
		}

		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}

}
