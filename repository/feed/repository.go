package feed

import (
	"../../errors"
	"../../models"
	"database/sql"
	_ "github.com/lib/pq"
)

type FeedRepositoryRealisation struct {
	feedDB *sql.DB
}

func NewFeedRepositoryRealisation(db *sql.DB) FeedRepositoryRealisation {
	return FeedRepositoryRealisation{feedDB: db}

}

func (Data FeedRepositoryRealisation) GetUserFeedById(id int, count int) (models.Feed, error) {
	rows, err := Data.feedDB.Query("select posts.txt_data, photos.url, photos.photo_likes_count, PhotosLikes.photo_was_like, posts.post_likes_count, posts.attachments, PostsLikes.post_was_like from (posts INNER JOIN feeds ON feeds.post_id=posts.post_id) INNER JOIN users ON users.u_id = feeds.u_id AND users.u_id = $1 LEFT JOIN photos ON photos.photo_id = posts.photo_id LEFT JOIN PhotosLikes ON PhotosLikes.photo_id = Posts.photo_id LEFT JOIN PostsLikes ON PostsLikes.post_id = Posts.post_id", id)
	if err != nil {
		return models.Feed{}, errors.FailReadFromDB
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
			return models.Feed{}, errors.FailReadToVar
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
	Feed := models.Feed{}
	Feed.Posts = posts

	return Feed, nil

}

func (Data FeedRepositoryRealisation) GetUserFeedByEmail(email string, count int) (models.Feed, error) {
	rows, err := Data.feedDB.Query("select posts.txt_data, photos.url, photos.photo_likes_count, PhotosLikes.photo_was_like, posts.post_likes_count, posts.attachments, PostsLikes.post_was_like from (posts INNER JOIN feeds ON feeds.post_id=posts.post_id) INNER JOIN users ON users.u_id = feeds.u_id AND users.mail = $1 LEFT JOIN photos ON photos.photo_id = posts.photo_id LEFT JOIN PhotosLikes ON PhotosLikes.photo_id = Posts.photo_id LEFT JOIN PostsLikes ON PostsLikes.post_id = Posts.post_id", email)
	if err != nil {
		return models.Feed{}, errors.FailReadFromDB
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
			return models.Feed{}, errors.FailReadToVar
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
	Feed := models.Feed{}
	Feed.Posts = posts

	return Feed, nil

}
