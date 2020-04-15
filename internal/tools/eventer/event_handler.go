package eventer

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"main/internal/message"
	"main/internal/models"
	"net"
)

type Eventer struct {
	userId    int
	messageDB message.MessageRepository
}

func NewEventer(user int, messages message.MessageRepository) Eventer {
	return Eventer{userId: user, messageDB: messages}
}

func (EV Eventer) WriteNewMessage(conn net.Conn) {
	req := wsutil.NewReader(conn, ws.StateClientSide)
	decoder := json.NewDecoder(req)

	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(resp)
	defer resp.Flush()

	answer := models.JSONEvent{}
	msg := new(models.Message)

	if err := decoder.Decode(msg); err != nil {
		answer.Event = "wrong json"
		encoder.Encode(&answer)
		fmt.Println("DECODE JSON ERROR : ", err)
		return
	}

	if err := EV.messageDB.AddNewMessage(EV.userId, *msg); err != nil {
		answer.Event = "can't add message"
		encoder.Encode(&answer)
		fmt.Println("ADD NEW MESSAGES ERROR : ", err)
		return
	}

	answer.Event = "ok"
	encoder.Encode(&answer)
}

func (EV Eventer) GetNewMessages(conn net.Conn) {
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(resp)
	answer := models.JSONEvent{}

	messages, err := EV.messageDB.ReceiveNewMessages(EV.userId)

	if err != nil {
		answer.Event = "can't get new messages"
		encoder.Encode(&answer)
		resp.Flush()
		fmt.Println("GET NEW MESSAGES ERROR : ", err)
		return
	}

	answer.Event = "new message"

	for iter, _ := range messages {
		answer.Message = messages[iter]
		encoder.Encode(&answer)
		resp.Flush()
	}

}
