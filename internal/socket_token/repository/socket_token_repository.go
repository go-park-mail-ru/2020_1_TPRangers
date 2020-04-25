package repository

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type TokenRepositoryRealisation struct {
	tokenDB *redis.Client
}

func NewTokenRepositoryRealisation(addr, pass string) TokenRepositoryRealisation {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return TokenRepositoryRealisation{tokenDB: client}
}

func (TR TokenRepositoryRealisation) AddNewToken(tokenValue string , userId int) error {
	return TR.tokenDB.Set(tokenValue , userId , time.Hour).Err()
}

func (TR TokenRepositoryRealisation) GetUserIdByToken(tokenValue string) (int , error) {
	sId, err := TR.tokenDB.Get(tokenValue).Result()
	resId, _ := strconv.Atoi(sId)

	if err != nil {
		resId = -1
	}

	return resId, err
}
