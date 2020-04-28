package tools

import "net"

type EventerInterface interface {
	GetNewMessages(net.Conn)
	WriteNewMessage(net.Conn)
}
