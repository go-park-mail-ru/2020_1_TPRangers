package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
	"main/internal/models"
	"math/rand"
	"testing"
)

func TestFeedRepositoryRealisation_GetUserPostsById(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 4
	rVal := uuid.NewV4()

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {
		uId := rand.Int()
		post := models.Post{
			Id:            rand.Int(),
			Text:          rVal.String(),
			Photo:         models.Photo{},
			Attachments:   "",
			Likes:         rand.Int(),
			WasLike:       false,
			Creation:      rVal.String(),
			AuthorName:    rVal.String(),
			AuthorSurname: rVal.String(),
			AuthorUrl:     rVal.String(),
			AuthorPhoto:   rVal.String(),
		}
		mock.ExpectBegin()
		mock.ExpectPrepare(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`)
		if errs[iter] == nil {
			mock.ExpectQuery(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`).
				WithArgs(uId).
				WillReturnRows(sqlmock.NewRows([]string{"P.post_id", "PH.photo_id", "P.txt_data", "P.posts_likes_count", "P.creation_date", "P.attachments", "UPL.postlike_id", "PH.url", "A.login", "A.name", "A.surname", "AP.url"}).AddRow(post.Id, post.Photo.Id, post.Text, post.Likes, "&postTime", post.Attachments, 2, post.Photo.Url, post.AuthorUrl, post.AuthorName, post.AuthorSurname, post.AuthorPhoto))

		} else {
			mock.ExpectQuery(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`).
				WithArgs(uId).WillReturnError(errs[iter])

		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		stmt, err := tx.Prepare("SELECT P.post_id,PH.photo_id,P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url,A.login,A.name,A.surname,AP.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = $1 AND P.post_id = UPL.post_id) LEFT JOIN PostsAuthor PA ON(PA.post_id=P.post_id) LEFT JOIN Users A ON(PA.u_id=A.u_id) INNER JOIN Photos AP ON(A.photo_id=AP.photo_id) WHERE UP.u_id = $1")
		var postTime string
		var likeId int
		err = stmt.QueryRow(uId).Scan(&post.Id, &post.Photo.Id, &post.Text, &post.Likes, &postTime, &post.Attachments, &likeId, &post.Photo.Url, &post.AuthorUrl, &post.AuthorName, &post.AuthorSurname, &post.AuthorPhoto)

		if err == errs[iter] {
			err = nil
		}
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}

}

func TestFeedRepositoryRealisation_GetUserPostsByLogin(t *testing.T) {

	db, mock, _ := sqlmock.New()
	testCounter := 4
	rVal := uuid.NewV4()

	errs := []error{nil, sql.ErrNoRows, nil, nil}

	for iter := 0; iter < testCounter; iter++ {
		uId := rand.Int()
		login := rVal.String()
		post := models.Post{
			Id:            rand.Int(),
			Text:          rVal.String(),
			Photo:         models.Photo{},
			Attachments:   "",
			Likes:         rand.Int(),
			WasLike:       false,
			Creation:      rVal.String(),
			AuthorName:    rVal.String(),
			AuthorSurname: rVal.String(),
			AuthorUrl:     rVal.String(),
			AuthorPhoto:   rVal.String(),
		}
		mock.ExpectBegin()

		mock.ExpectPrepare(`SELECT u_id FROM Users WHERE login \= \$1`)
		mock.ExpectQuery(`SELECT u_id FROM Users WHERE login \= \$1`).WithArgs(login).
			WillReturnRows(sqlmock.NewRows([]string{"u_id"}).AddRow(uId))

		mock.ExpectPrepare(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`)
		if errs[iter] == nil {
			mock.ExpectQuery(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`).
				WithArgs(uId).
				WillReturnRows(sqlmock.NewRows([]string{"P.post_id", "PH.photo_id", "P.txt_data", "P.posts_likes_count", "P.creation_date", "P.attachments", "UPL.postlike_id", "PH.url", "A.login", "A.name", "A.surname", "AP.url"}).AddRow(post.Id, post.Photo.Id, post.Text, post.Likes, "&postTime", post.Attachments, 2, post.Photo.Url, post.AuthorUrl, post.AuthorName, post.AuthorSurname, post.AuthorPhoto))
		} else {
			mock.ExpectQuery(`SELECT P\.post_id,PH\.photo_id,P\.txt_data, P\.posts_likes_count, P\.creation_date,P\.attachments,UPL\.postlike_id,PH\.url,A\.login,A\.name,A\.surname,AP\.url FROM UsersPosts UP INNER JOIN Posts P ON\(P\.post_id\=UP\.post_id\) LEFT JOIN Photos PH ON\(PH\.photo_id\=P\.photo_id\) LEFT JOIN UsersPostsLikes UPL ON\(UPL\.u_id \= \$1 AND P\.post_id \= UPL\.post_id\) LEFT JOIN PostsAuthor PA ON\(PA\.post_id\=P\.post_id\) LEFT JOIN Users A ON\(PA\.u_id\=A\.u_id\) INNER JOIN Photos AP ON\(A\.photo_id\=AP\.photo_id\) WHERE UP\.u_id \= \$1`).
				WithArgs(uId).WillReturnError(errs[iter])
		}
		mock.ExpectCommit()

		tx, err := db.Begin()

		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		smtId, err := tx.Prepare("SELECT u_id FROM Users WHERE login = $1")
		if err != nil {
			fmt.Println("db not created", err)
			return
		}
		currentUId := 0
		err = smtId.QueryRow(login).Scan(&currentUId)
		if err != nil {
			fmt.Println("db not created", err)
			return
		}

		stmt, err := tx.Prepare("SELECT P.post_id,PH.photo_id,P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url,A.login,A.name,A.surname,AP.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = $1 AND P.post_id = UPL.post_id) LEFT JOIN PostsAuthor PA ON(PA.post_id=P.post_id) LEFT JOIN Users A ON(PA.u_id=A.u_id) INNER JOIN Photos AP ON(A.photo_id=AP.photo_id) WHERE UP.u_id = $1")

		var postTime string
		var likeId int
		err = stmt.QueryRow(currentUId).Scan(&post.Id, &post.Photo.Id, &post.Text, &post.Likes, &postTime, &post.Attachments, &likeId, &post.Photo.Url, &post.AuthorUrl, &post.AuthorName, &post.AuthorSurname, &post.AuthorPhoto)

		if err == errs[iter] {
			err = nil
		}
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}

	}
}
