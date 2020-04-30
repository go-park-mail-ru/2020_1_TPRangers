package repository

import (
	"database/sql"
	"errors"
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

	chat := int64(0)
	groupType := "gch"
	err := errors.New("")

	if message.ChatId[:1] != "c" {
		chat, err = strconv.ParseInt(message.ChatId, 10, 64)

		if err != nil {
			return nil
		}
		groupType = "pch"
	} else {
		chat, err = strconv.ParseInt(message.ChatId[1:], 10, 64)

		if err != nil {
			return nil
		}
		groupType = "gch"
	}
	msgRow := MR.messageDB.QueryRow("INSERT INTO Messages ("+groupType+"_id,u_id,txt,send_time) VALUES($1,$2,$3,$4) RETURNING msg_id", chat, author, message.Text, time.Now())

	msgId := 0
	err = msgRow.Scan(&msgId)

	if err != nil {
		return err
	}

	recRows, err := MR.messageDB.Query("SELECT u_id FROM ChatsUsers WHERE "+groupType+"_id = $1", chat)
	defer func() {
		if recRows != nil {
			recRows.Close()
		}
	}()

	if err != nil {
		return err
	}

	for recRows.Next() {
		reciever := 0

		err = recRows.Scan(&reciever)
		_, err = MR.messageDB.Exec("INSERT INTO NewMessages (msg_id,receiver_id) VALUES($1,$2)", msgId, reciever)

	}

	return nil
}

func (MR MessageRepositoryRealisation) ReceiveNewMessages(userId int) ([]models.Message, error) {

	msgsArray := make([]models.Message, 0)

	msgsRow, err := MR.messageDB.Query("SELECT M.msg_id, M.gch_id, M.pch_id ,M.u_id ,M.send_time ,M.txt , U.name,P.url FROM Messages M "+
		"INNER JOIN NewMessages NM ON(NM.msg_id=M.msg_id) LEFT JOIN GroupChats GC ON(M.gch_id=GC.ch_id) INNER JOIN Users U ON(U.u_id=M.u_id) "+
		"LEFT JOIN Photos P ON CASE "+
		"WHEN P.photo_id=GC.photo_id AND M.pch_id = 0 THEN 1 "+
		"WHEN U.photo_id=P.photo_id AND M.gch_id = 0 THEN 1 "+
		"ELSE 0 "+
		"END = 1 "+
		"WHERE NM.receiver_id = $1 AND M.del_stat = true", userId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for msgsRow.Next() {

		msg := new(models.Message)
		msgId := 0
		userid := 0

		var isPrivate *int64
		var isGroup *int64
		err = msgsRow.Scan(&msgId, &isGroup, &isPrivate, &userid, &msg.Time, &msg.Text, &msg.ChatName, &msg.ChatPhoto)


		if *isGroup != int64(0) {
			msg.ChatId = "c" + strconv.FormatInt(*isGroup,10)
		} else {
			if *isPrivate != int64(0) {
				msg.ChatId = strconv.FormatInt(*isPrivate,10)
			}
		}

		if err != nil {
			return nil, err
		}

		MR.messageDB.Exec("DELETE FROM NewMessages WHERE msg_id = $1 AND receiver_id = $2", msgId, userId)

		msgsArray = append(msgsArray, *msg)

	}

	return msgsArray, nil
}
