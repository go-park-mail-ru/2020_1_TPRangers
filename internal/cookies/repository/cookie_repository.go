package repository

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type CookieRepositoryRealisation struct {
	sessionDB *redis.Client
}

func NewCookieRepositoryRealisation(addr, pass string) CookieRepositoryRealisation {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       0,    // use default DB
	})

	return CookieRepositoryRealisation{sessionDB: client}
}

func (cookR CookieRepositoryRealisation) AddCookie(id int, cookieValue string, exprTime time.Duration) error {

	err := cookR.sessionDB.Set(cookieValue, id, exprTime).Err()

	return err
}

func (cookR CookieRepositoryRealisation) DeleteCookie(cookieValue string) error {

	err := cookR.sessionDB.Del(cookieValue).Err()

	return err

}

func (cookR CookieRepositoryRealisation) GetUserIdByCookie(cookieValue string) (int, error) {

	sId, err := cookR.sessionDB.Get(cookieValue).Result()
	resId, _ := strconv.Atoi(sId)

	return resId, err
}