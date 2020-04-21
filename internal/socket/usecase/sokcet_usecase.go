package usecase

import (
	"main/internal/message"
	stocken "main/internal/socket_token"
	cr "main/internal/tools/connection_reciever"
	"main/internal/tools/eventer"
	"net"
)

type SocketUseCaseRealisation struct {
	messageDB message.MessageRepository
	tokenDB   stocken.TokenRepository
}

func NewSocketUseCaseRealisation(messageDB message.MessageRepository, tokenDB stocken.TokenRepository) SocketUseCaseRealisation {
	return SocketUseCaseRealisation{messageDB: messageDB, tokenDB: tokenDB}
}

func (SU SocketUseCaseRealisation) CheckToken(tokenValue string) (int, error) {
	return SU.tokenDB.GetUserIdByToken(tokenValue)
}

func (SU SocketUseCaseRealisation) AddToConnectionPool(conn net.Conn, userId int) error {

	eventer := eventer.NewEventer(userId, SU.messageDB)

	reciever, err := cr.NewConnReciever(conn, eventer)
	if err == nil {
		reciever.StartRecieving()
	}
	return err
}
