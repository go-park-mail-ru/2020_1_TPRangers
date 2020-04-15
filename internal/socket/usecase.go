package socket

import "net"

type SocketUseCase interface {
	AddToConnectionPool(net.Conn, int) error
}
