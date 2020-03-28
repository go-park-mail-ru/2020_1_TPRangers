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



func (Like LikeRepositoryRealisation) LikePhoto