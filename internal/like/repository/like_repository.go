package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type LikeRepositoryRealisation struct {
	likeDB *sql.DB
}

func NewLikeRepositoryRealisation(db *sql.DB) LikeRepositoryRealisation {
	return LikeRepositoryRealisation{likeDB: db}
}



func (Like LikeRepositoryRealisation) LikePhoto(photoId, userId int) error {
	Like.likeDB.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count + 1 WHERE	photo_id =$1", photoId)

	like_id := int64(0)

	row := Like.likeDB.QueryRow("INSERT INTO UsersPhotosLikes (u_id,photo_id) VALUES ($1,$2) RETURNING photolike_id",userId , photoId)

	err := row.Scan(&like_id)

	return err
}

func (Like LikeRepositoryRealisation) DislikePhoto(photoId, userId int) error {
	Like.likeDB.Exec("UPDATE Photos SET photos_likes_count = photos_likes_count - 1 WHERE photo_id = $1", photoId)

	like_id := int64(0)

	row := Like.likeDB.QueryRow("DELETE FROM UserPhotosLike WHERE u_id = $1 AND photo_id = $2 RETURNING photolike_id",userId , photoId)

	err := row.Scan(&like_id)

	return err
}

