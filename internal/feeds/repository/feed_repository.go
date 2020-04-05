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
	rows, err := Data.feedDB.Query("select posts.txt_data, photos.url, photos.photos_likes_count, PhotosLikes.photo_was_like, posts.posts_likes_count, posts.attachments, PostsLikes.post_was_like from (posts INNER JOIN feeds ON feeds.post_id=posts.post_id) INNER JOIN users ON users.u_id = feeds.u_id AND users.u_id = $1 LEFT JOIN photos ON photos.photo_id = posts.photo_id LEFT JOIN PhotosLikes ON PhotosLikes.photo_id = Posts.photo_id LEFT JOIN PostsLikes ON PostsLikes.post_id = Posts.post_id", id)
	if err != nil {

		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ id")
		}

		return nil, errors.FailReadFromDB
	}
	posts := []models.Post{}
	var photoUrl interface{}
	var photowasLike interface{}
	var photoLikes interface{}
	var postAttachments interface{}
	var postWasLike interface{}
	var postLikes interface{}
	var postText interface{}
	var i int
	for rows.Next() {
		if i > count {
			break
		}
		post := models.Post{}
		err := rows.Scan(&postText, &photoUrl, &photoLikes, &photowasLike, &postLikes, &postAttachments, &postWasLike)

		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ id")
		}

		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ id")
			return nil, errors.FailReadToVar
		}

		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if photowasLike == nil || photowasLike.(bool) == false {
			post.Photo.WasLike = false
		} else {
			post.Photo.WasLike = true
		}
		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if postAttachments == nil {
			post.Attachments = ""
		} else {
			post.Attachments = postAttachments.(string)
		}
		if postWasLike == nil || postWasLike.(bool) == false {
			post.WasLike = false
		} else {
			post.WasLike = true
		}
		if postLikes == nil {
			post.Likes = 0
		} else {
			post.Likes = int(postLikes.(int64))
		}
		if postText == nil {
			post.Text = ""
		} else {
			post.Text = postText.(string)
		}

		posts = append(posts, post)
		i++
	}

	return posts, nil

}

func (Data FeedRepositoryRealisation) GetUserFeedByEmail(email string, count int) ([]models.Post, error) {
	rows, err := Data.feedDB.Query("select posts.txt_data, photos.url, photos.photos_likes_count, PhotosLikes.photo_was_like, posts.posts_likes_count, posts.attachments, PostsLikes.post_was_like from (posts INNER JOIN feeds ON feeds.post_id=posts.post_id) INNER JOIN users ON users.u_id = feeds.u_id AND users.mail = $1 LEFT JOIN photos ON photos.photo_id = posts.photo_id LEFT JOIN PhotosLikes ON PhotosLikes.photo_id = Posts.photo_id LEFT JOIN PostsLikes ON PostsLikes.post_id = Posts.post_id", email)
	if err != nil {

		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ id")
		}

		return nil, errors.FailReadFromDB
	}
	posts := []models.Post{}
	var photoUrl interface{}
	var photowasLike interface{}
	var photoLikes interface{}
	var postAttachments interface{}
	var postWasLike interface{}
	var postLikes interface{}
	var postText interface{}
	var i int
	for rows.Next() {
		if i > count {
			break
		}
		post := models.Post{}
		err := rows.Scan(&postText, &photoUrl, &photoLikes, &photowasLike, &postLikes, &postAttachments, &postWasLike)

		if err != nil {
			fmt.Println(err.Error(), "FEED ================================ EMAIL")
		}

		if err != nil {
			return nil, errors.FailReadToVar
		}

		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if photowasLike == nil || photowasLike.(bool) == false {
			post.Photo.WasLike = false
		} else {
			post.Photo.WasLike = true
		}
		if photoLikes == nil {
			post.Photo.Likes = 0
		} else {
			post.Photo.Likes = int(photoLikes.(int64))
		}
		if postAttachments == nil {
			post.Attachments = ""
		} else {
			post.Attachments = postAttachments.(string)
		}
		if postWasLike == nil || postWasLike.(bool) == false {
			post.WasLike = false
		} else {
			post.WasLike = true
		}
		if postLikes == nil {
			post.Likes = 0
		} else {
			post.Likes = int(postLikes.(int64))
		}
		if postText == nil {
			post.Text = ""
		} else {
			post.Text = postText.(string)
		}

		posts = append(posts, post)
		i++
	}

	return posts, nil

}

func (Data FeedRepositoryRealisation) GetUserPostsById(id int) ([]models.Post, error) {

	feed := make([]models.Post, 0)

	row, err := Data.feedDB.Query("SELECT P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = $1 AND P.post_id = UPL.post_id) WHERE UP.u_id = $1", id)
	if err != nil {
		fmt.Println(err, "USER POSTS ERROR")
		return feed, err
	}

	for row.Next() {
		post := new(models.Post)
		var likeId sql.NullInt32
		likeId.Scan(-1)

		err = row.Scan(&post.Text, &post.Likes, &post.Creation, &post.Attachments, &likeId, &post.Photo.Url)

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

	feed := make([]models.Post, 0)

	userRow := Data.feedDB.QueryRow("SELECT u_id FROM Users WHERE login = $1",login)

	userId := 0
	userRow.Scan(&userId)

	row, err := Data.feedDB.Query("SELECT P.txt_data, P.posts_likes_count, P.creation_date,P.attachments,UPL.postlike_id,PH.url FROM UsersPosts UP INNER JOIN Posts P ON(P.post_id=UP.post_id) LEFT JOIN Photos PH ON(PH.photo_id=P.photo_id) LEFT JOIN UsersPostsLikes UPL ON(UPL.u_id = UP.u_id AND P.post_id = UPL.post_id) WHERE UP.u_id = $1", userId)
	if err != nil {
		fmt.Println(err, "USER POSTS ERROR")
		return feed, err
	}

	for row.Next() {
		post := new(models.Post)

		var likeId sql.NullInt32
		likeId.Scan(-1)

		err = row.Scan(&post.Text, &post.Likes, &post.Creation, &post.Attachments, &likeId, &post.Photo.Url)

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

	if newPost.Photo.Url != "" {
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
	if err != nil {
		fmt.Println(err , "ERROR ON ADDING NEW POST TO DATABASE")
		return errors.FailSendToDB
	}

	Data.feedDB.Exec("INSERT INTO UsersPosts (u_id,post_id) VALUES($1,$2)",uId,postId)

	return nil
}
