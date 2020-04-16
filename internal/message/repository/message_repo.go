package repository

import (
	"database/sql"
	"github.com/go-redis/redis"
	"main/internal/models"
)

type MessageRepositoryRealisation struct {
	msgNotifier *redis.Client
	messageDB   *sql.DB
}

func NewMessageRepositoryRealisation(addr, pass string , db *sql.DB) MessageRepositoryRealisation {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	return MessageRepositoryRealisation{msgNotifier: client , messageDB: db}
}

func (MR MessageRepositoryRealisation) AddNewMessage(author int, message models.Message) error {
	 // some db insert

	reciever := message.Receiver

	err := MR.msgNotifier
}
