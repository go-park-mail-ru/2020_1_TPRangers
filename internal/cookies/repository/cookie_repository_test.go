package repository

import (
	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"
)

func newTestRedis() (*redismock.ClientMock, *redis.Client) {
	mr, _ := miniredis.Run()

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client), client
}

func TestCookieRepositoryRealisation_DeleteCookie(t *testing.T) {
	_, client := newTestRedis()

	expectedBehaviour := []error{nil, nil}

	cRepo := CookieRepositoryRealisation{sessionDB: client}

	for iter, _ := range expectedBehaviour {

		c := uuid.NewV4()
		cVal := c.String()

		client.Set(cVal, "123", 5*time.Millisecond)

		if err := cRepo.DeleteCookie(cVal); err != expectedBehaviour[iter] {
			t.Error(iter, err, expectedBehaviour[iter])
		}

	}

}

func TestCookieRepositoryRealisation_AddCookie(t *testing.T) {
	_, client := newTestRedis()

	expectedBehaviour := []error{nil, nil}

	cRepo := CookieRepositoryRealisation{sessionDB: client}

	for iter, _ := range expectedBehaviour {

		c := uuid.NewV4()
		cVal := c.String()

		client.Set(cVal, 1, 5*time.Millisecond)

		if err := cRepo.AddCookie(1, cVal, 5*time.Millisecond); err != expectedBehaviour[iter] {
			t.Error(iter, err, expectedBehaviour[iter])
		}

	}

}

func TestCookieRepositoryRealisation_GetUserIdByCookie(t *testing.T) {
	_, client := newTestRedis()

	expectedBehaviour := []error{nil, nil}

	cRepo := CookieRepositoryRealisation{sessionDB: client}

	for iter, _ := range expectedBehaviour {

		c := uuid.NewV4()
		cVal := c.String()

		client.Set(cVal, 1, 5*time.Millisecond)

		if _, err := cRepo.GetUserIdByCookie(cVal); err != expectedBehaviour[iter] {
			t.Error(iter, err, expectedBehaviour[iter])
		}

	}

}
