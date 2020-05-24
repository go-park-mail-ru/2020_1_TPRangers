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
	var groupType string
	err := errors.New("")
	if err != nil {
		return err
	}

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

	var msgRow *sql.Row
	if message.Sticker == "" {
		msgRow = MR.messageDB.QueryRow("INSERT INTO Messages ("+groupType+"_id,u_id,txt,send_time,attach_link) VALUES($1,$2,$3,$4,$5) RETURNING msg_id", chat, author, message.Text, time.Now() , message.Attachment)
	} else {
		msgRow = MR.messageDB.QueryRow("INSERT INTO Messages ("+groupType+"_id,u_id,txt,send_time,sticker_link,attach_link) VALUES($1,$2,$3,$4,$5,$6) RETURNING msg_id", chat, author, message.Text, time.Now(), message.Sticker , message.Attachment)
	}

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
		if err != nil {
			return err
		}
		_, err = MR.messageDB.Exec("INSERT INTO NewMessages (msg_id,receiver_id) VALUES($1,$2)", msgId, reciever)
		if err != nil {
			return err
		}

	}

	return nil
}

func (MR MessageRepositoryRealisation) ReceiveNewMessages(userId int) ([]models.Message, error) {

	msgsArray := make([]models.Message, 0)

	msgsRow, err := MR.messageDB.Query("SELECT M.msg_id, M.gch_id, M.pch_id ,M.u_id ,M.send_time ,M.txt , U.name,P.url, M.sticker_link , M.attach_link FROM Messages M "+
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
		var stickerLink *string
		err = msgsRow.Scan(&msgId, &isGroup, &isPrivate, &userid, &msg.Time, &msg.Text, &msg.ChatName, &msg.ChatPhoto, &stickerLink , &msg.Attachment)

		if *isGroup != int64(0) {
			msg.ChatId = "c" + strconv.FormatInt(*isGroup, 10)
		} else {
			if *isPrivate != int64(0) {
				msg.ChatId = strconv.FormatInt(*isPrivate, 10)
			}
		}

		if stickerLink != nil {
			msg.Sticker = *stickerLink
		}

		if err != nil {
			return nil, err
		}

		_, err = MR.messageDB.Exec("DELETE FROM NewMessages WHERE msg_id = $1 AND receiver_id = $2", msgId, userId)
		if err != nil {
			return nil, err
		}

		msgsArray = append(msgsArray, *msg)

	}

	return msgsArray, nil
}
