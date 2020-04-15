package usecase

import (
	"main/internal/message"
	cr "main/internal/tools/connection_reciever"
	"main/internal/tools/eventer"
	"net"
)

type SocketUseCaseRealisation struct {
	messageRepo message.MessageRepository
}

func NewSocketUseCaseRealisation(mR message.MessageRepository) SocketUseCaseRealisation {
	return SocketUseCaseRealisation{messageRepo: mR}
}

func (SU SocketUseCaseRealisation) AddToConnectionPool(conn net.Conn, userId int) error {

	eventer := eventer.NewEventer(userId, SU.messageRepo)

	reciever, err := cr.NewConnReciever(conn, eventer)
	if err == nil {
		reciever.StartRecieving()
	}
	return err
}
