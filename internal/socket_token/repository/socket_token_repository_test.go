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

func TestTokenRepositoryRealisation_AddNewToken(t *testing.T) {
	_, client := newTestRedis()

	expectedBehaviour := []error{nil, nil}

	cRepo := TokenRepositoryRealisation{tokenDB: client}

	for iter, _ := range expectedBehaviour {

		c := uuid.NewV4()
		cVal := c.String()

		client.Set(cVal, "123", 5*time.Millisecond)

		if err := cRepo.AddNewToken("123", 1); err != expectedBehaviour[iter] {
			t.Error(iter, err, expectedBehaviour[iter])
		}

	}
}

func TestTokenRepositoryRealisation_GetUserIdByToken(t *testing.T) {
	_, client := newTestRedis()

	expectedBehaviour := []error{nil, nil}

	cRepo := TokenRepositoryRealisation{tokenDB: client}

	for iter, _ := range expectedBehaviour {

		c := uuid.NewV4()
		tVal := c.String()

		client.Set(tVal, "123", 5*time.Millisecond)

		if _, err := cRepo.GetUserIdByToken(tVal); err != expectedBehaviour[iter] {
			t.Error(iter, err, expectedBehaviour[iter])
		}

	}

	c := uuid.NewV4()
	tVal := c.String()

	if _, err := cRepo.GetUserIdByToken(tVal); err == nil {
		t.Error("ERROR")
	}

}
