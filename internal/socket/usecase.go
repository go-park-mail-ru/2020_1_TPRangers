package socket

import "net"

type SocketUseCase interface {
	CheckToken(string) (int, error)
	AddToConnectionPool(net.Conn, int) error
}
