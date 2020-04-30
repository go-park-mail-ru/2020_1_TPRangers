package eventer

import (
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"main/internal/message"
	"main/models"
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

	fmt.Println("wrote new message")

	req := wsutil.NewReader(conn, ws.StateServerSide)
	decoder := json.NewDecoder(req)
	hdr, err := req.NextFrame()
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)

	if err != nil || hdr.OpCode == ws.OpClose {
		fmt.Println("connection closed")
		conn.Close()
		return
	}

	defer func() {
		resp.Flush()
	}()

	msg := new(models.Message)

	if err = decoder.Decode(&msg); err != nil {

		fmt.Println("DECODE JSON ERROR : ", err)
		return
	}

	if err := EV.messageDB.AddNewMessage(EV.userId, *msg); err != nil {

		fmt.Println("ADD NEW MESSAGES ERROR : ", err)
		return
	}

}

func (EV Eventer) GetNewMessages(conn net.Conn) {
	resp := wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
	encoder := json.NewEncoder(resp)
	answer := models.JSONEvent{}

	defer func() {
		resp.Flush()
	}()
	messages, err := EV.messageDB.ReceiveNewMessages(EV.userId)

	if err != nil {
		answer.Event = "can't get new messages"
		encoder.Encode(&answer)
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
