package repository

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
	"main/internal/models"
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

	msgRow := MR.messageDB.QueryRow("INSERT INTO Messages (ch_id,u_id,txt,send_time) VALUES($1,$2,$3,$4) RETURNING msg_id", message.ChatId, author, message.Text, time.Now())

	msgId := 0
	err := msgRow.Scan(&msgId)

	if err != nil {
		return err
	}

	recRows, err := MR.messageDB.Query("SELECT u_id FROM ChatsUsers WHERE ch_id = $1", message.ChatId)
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
		_ , err = MR.messageDB.Exec("INSERT INTO NewMessages (msg_id,receiver_id) VALUES($1,$2)", msgId, reciever)

	}

	return nil
}

func (MR MessageRepositoryRealisation) ReceiveNewMessages(userId int) ([]models.Message, error) {

	msgsArray := make([]models.Message, 0)

	msgsRow, err := MR.messageDB.Query("SELECT M.msg_id ,M.ch_id , M.u_id ,  M.send_time , M.txt  FROM Messages M INNER JOIN NewMessages NM ON(NM.msg_id=M.msg_id) WHERE NM.receiver_id = $1 AND M.del_stat = true", userId)

	if err != nil {
		return nil, err
	}

	for msgsRow.Next() {

		msg := new(models.Message)
		msgId := 0
		userid := 0
		err = msgsRow.Scan(&msgId, &msg.ChatId, &userid, &msg.Time, &msg.Text)

		if err != nil {
			return nil, err
		}

		MR.messageDB.Exec("DELETE FROM NewMessages WHERE msg_id = $1 AND receiver_id = $2", msgId, userId)

		msgsArray = append(msgsArray, *msg)

	}

	return msgsArray, nil
}
