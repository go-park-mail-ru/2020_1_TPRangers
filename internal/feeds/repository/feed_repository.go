package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"main/internal/models"
	"main/internal/tools/errors"
	"time"
)

type FeedRepositoryRealisation struct {
	feedDB *sql.DB
}

func NewFeedRepositoryRealisation(db *sql.DB) FeedRepositoryRealisation {
	return FeedRepositoryRealisation{feedDB: db}

}

func (Data FeedRepositoryRealisation) GetUserFeedById(id int, count int) ([]models.Post, error) {
	rows, err := Data.feedDB.Query("select  p.post_id, p.txt_data, p.attachments, p.posts_likes_count, p.creation_date, ph.photo_id, ph.url, ph.photos_likes_count, u.name, u.surname, u.login " +
	"from posts AS p INNER JOIN postsauthor AS pa ON (p.post_id = pa.post_id) INNER JOIN users AS u ON (pa.u_id = u.u_id) LEFT JOIN photos AS ph ON p.photo_id = ph.photo_id;")
	if err != nil {
		return nil, errors.FailReadFromDB
	}
	posts := []models.Post{}
	i := 0
	for rows.Next() {
		if i > count {
			break
		}
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.Text, &post.Attachments, &post.Likes, &post.Creation, &post.Photo.Id, &post.Photo.Url, &post.Photo.Likes, &post.AuthorName, &post.AuthorSurname, &post.AuthorUrl)

		additional_row := Data.feedDB.QueryRow("select upl.postlike_id, uphl.photolike_id, ph.url from userspostslikes AS upl RIGHT JOIN posts AS p " +
			"ON (p.post_id = upl.post_id) LEFT JOIN usersphotoslikes AS uphl ON (p.photo_id = uphl.photo_id) INNER JOIN postsauthor AS pa ON (pa.post_id = p.post_id) " +
			"LEFT JOIN users AS u ON (u.u_id = pa.u_id) LEFT JOIN photos AS ph ON (u.photo_id = ph.photo_id) WHERE p.post_id = $1;", post.Id)
		var postLikes *int
		var photoLikes *int
		additional_row.Scan(&postLikes, &photoLikes, &post.AuthorPhoto)
		if postLikes != nil {
			post.WasLike = true
		} else {
			post.WasLike = false
		}
		if photoLikes != nil {
			post.Photo.WasLike = true
		} else {
			post.Photo.WasLike = false
		}
		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ id")
			return nil, errors.FailReadToVar
		}

		posts = append(posts, post)
		i++
	}

	return posts, nil

}

func (Data FeedRepositoryRealisation) GetUserPostsById(id int) ([]models.Post, error) {

	feed := make([]models.Post, 0)
	row, err := Data.feedDB.Query("SELECT P.post_id,PH.photo_id,P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url,A.login,A.name,A.surname,AP.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = $1 AND P.post_id = UPL.post_id) LEFT JOIN PostsAuthor PA ON(PA.post_id=P.post_id) LEFT JOIN Users A ON(PA.u_id=A.u_id) INNER JOIN Photos AP ON(A.photo_id=AP.photo_id) WHERE UP.u_id = $1", id)
	if err != nil {
		fmt.Println(err, "USER POSTS ERROR")
		return feed, err
	}

	for row.Next() {
		post := new(models.Post)
		var likeId sql.NullInt32
		var postTime time.Time
		err = row.Scan(&post.Id,&post.Photo.Id,&post.Text, &post.Likes, &postTime, &post.Attachments, &likeId, &post.Photo.Url , &post.AuthorUrl, &post.AuthorName, &post.AuthorSurname, &post.AuthorPhoto)

		post.Creation = postTime.Format("2006-01-02 15:04:05")

		if err != nil {
			fmt.Println(err.Error(), "USER POSTS BY ID")
		}

		value , _ := likeId.Value()

		if value == nil {
			post.WasLike = false
		} else {
			post.WasLike = true
		}

		feed = append(feed , *post)
	}

	return feed, nil

}

func (Data FeedRepositoryRealisation) GetUserPostsByLogin(login string) ([]models.Post, error) {

	userRow := Data.feedDB.QueryRow("SELECT u_id FROM Users WHERE login = $1",login)

	userId := 0
	userRow.Scan(&userId)


	feed := make([]models.Post, 0)
	row, err := Data.feedDB.Query("SELECT P.post_id,PH.photo_id,P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,PH.url,A.login,A.name,A.surname,AP.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN PostsAuthor PA ON(PA.post_id=P.post_id) LEFT JOIN Users A ON(PA.u_id=A.u_id) INNER JOIN Photos AP ON(A.photo_id=AP.photo_id) WHERE UP.u_id = $1", userId)
	if err != nil {
		fmt.Println(err, "USER POSTS ERROR")
		return feed, err
	}

	for row.Next() {
		post := new(models.Post)
		var postTime time.Time
		err = row.Scan(&post.Id,&post.Photo.Id,&post.Text, &post.Likes, &postTime, &post.Attachments, &post.Photo.Url , &post.AuthorUrl, &post.AuthorName, &post.AuthorSurname, &post.AuthorPhoto)

		post.Creation = postTime.Format("2006-01-02 15:04:05")

		if err != nil {
			fmt.Println(err.Error(), "USER POSTS BY ID")
		}

		feed = append(feed , *post)
	}

	return feed, nil

}

func (Data FeedRepositoryRealisation) GetPostsOfOtherUserWhileLogged(login string, currentUserId int) ([]models.Post, error) {

	userRow := Data.feedDB.QueryRow("SELECT u_id FROM Users WHERE login = $1",login)

	userId := 0
	userRow.Scan(&userId)


	feed := make([]models.Post, 0)
	row, err := Data.feedDB.Query("SELECT P.post_id,PH.photo_id,P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url,A.login,A.name,A.surname,AP.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = $2 AND P.post_id = UPL.post_id) LEFT JOIN PostsAuthor PA ON(PA.post_id=P.post_id) LEFT JOIN Users A ON(PA.u_id=A.u_id) INNER JOIN Photos AP ON(A.photo_id=AP.photo_id) WHERE UP.u_id = $1", userId, currentUserId)
	if err != nil {
		fmt.Println(err, "USER POSTS ERROR")
		return feed, err
	}

	for row.Next() {
		post := new(models.Post)
		var likeId sql.NullInt32
		var postTime time.Time
		err = row.Scan(&post.Id,&post.Photo.Id,&post.Text, &post.Likes, &postTime, &post.Attachments, &likeId, &post.Photo.Url , &post.AuthorUrl, &post.AuthorName, &post.AuthorSurname, &post.AuthorPhoto)

		post.Creation = postTime.Format("2006-01-02 15:04:05")

		if err != nil {
			fmt.Println(err.Error(), "USER POSTS BY ID")
		}

		value , _ := likeId.Value()

		if value == nil {
			post.WasLike = false
		} else {
			post.WasLike = true
		}

		feed = append(feed , *post)
	}

	return feed, nil

}

func (Data FeedRepositoryRealisation) CreatePost(uId int,newPost models.Post) error {

	photo_id := 0

	if newPost.Photo.Url != nil {
		row := Data.feedDB.QueryRow("INSERT INTO photos (url, photos_likes_count) VALUES ($1 , 0) RETURNING photo_id", newPost.Photo.Url)

		errScan := row.Scan(&photo_id)

		if errScan != nil {
			fmt.Println(errScan,"ERR ON CREATE PHOTO FOR NEW POST")
		}
	}

	postRow , err := Data.feedDB.Query("INSERT INTO Posts (txt_data,photo_id,posts_likes_count,creation_date, attachments) VALUES($1 , $2 , $3 , $4 , $5) RETURNING post_id", newPost.Text, photo_id, 0 , time.Now(), newPost.Attachments)

	if err != nil {
		fmt.Println(err , "ERROR ON ADDING NEW POST TO DATABASE")
		return errors.FailSendToDB
	}

	postRow.Next()

	postId := 0
	err = postRow.Scan(&postId)
	if err == nil {
		Data.feedDB.Exec("INSERT INTO UsersPosts (u_id,post_id) VALUES($1,$2)",uId,postId)
		Data.feedDB.Exec("INSERT INTO PostsAuthor (u_id,post_id) VALUES($1,$2)",uId,postId)
		return nil

	}


	fmt.Println(err , "ERROR ON ADDING NEW POST TO DATABASE")
	return errors.FailSendToDB

}
