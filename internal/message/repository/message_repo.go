package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
	"main/internal/models"
	"strconv"
	"time"
)

type MessageRepositoryRealisation struct {
	msgNotifier *rejson.Handler
	messageDB   *sql.DB
}

func NewMessageRepositoryRealisation(addr, pass string, db *sql.DB) MessageRepositoryRealisation {

	rh := rejson.NewReJSONHandler()
	client, _ := redis.Dial("tcp", addr)
	rh.SetRedigoClient(client)

	return MessageRepositoryRealisation{msgNotifier: rh, messageDB: db}
}

func (MR MessageRepositoryRealisation) AddNewMessage(author int, message models.Message) error {

	_, err := MR.messageDB.Exec("INSERT INTO Messages (ch_id,u_id,txt,send_time) VALUES($1,$2,$3,$4)", message.ChatId, author, message.Text, time.Now())

	if err != nil {
		return err
	}

	recRows, err := MR.messageDB.Query("SELECT u_id FROM ChatsUsers WHERE ch_id = $1 AND u_id != $2", message.ChatId, author)
	defer func() {
		if recRows != nil {
			recRows.Close()
		}
	}()

	for recRows.Next() {
		reciever := 0

		err = recRows.Scan(&reciever)

		if err != nil {
			return err
		}

		var res interface{}
		if _, err = redis.Bytes(MR.msgNotifier.JSONGet(strconv.Itoa(author), "")); err != nil {
			res, err = MR.msgNotifier.JSONSet(strconv.Itoa(author), "", message)
		} else {
			MR.msgNotifier.JSONArrAppend(strconv.Itoa(author), "", message)
		}

		if err != nil || res.(string) != "OK" {
			return err
		}
	}
	return nil
}

func (MR MessageRepositoryRealisation) ReceiveNewMessages(userId int) ([]models.Message, error) {

	msgsInt, err := MR.msgNotifier.JSONArrLen(strconv.Itoa(userId), "")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	msgsArray := msgsInt.([]models.Message)

	for iter := msgsInt.(int); iter >= 0; iter-- {
		msg, err := MR.msgNotifier.JSONArrPop(strconv.Itoa(userId), "", iter)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		msgJSON := new(models.Message)

		err = json.Unmarshal(msg.([]byte), &msgJSON)

		if err != nil {
			fmt.Println(err)
			msgsArray = append(msgsArray, msg.(models.Message))
		} else {
			msgsArray = append(msgsArray, *msgJSON)
		}

	}

	return msgsArray, nil
}
