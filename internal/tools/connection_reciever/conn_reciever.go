package connection_reciever

import (
	"errors"
	"fmt"
	"github.com/mailru/easygo/netpoll"
	"main/internal/tools"
	"main/internal/tools/gpool"
	"net"
	"time"
)

type ConnReceiver struct {
	workPool   tools.GoPoolInterface
	netPool    netpoll.Poller
	connection net.Conn
	handler    tools.EventerInterface
}

func NewConnReciever(conn net.Conn, handler tools.EventerInterface) (ConnReceiver, error) {
	poller, err := netpoll.New(nil)

	return ConnReceiver{netPool: poller, workPool: gpool.New(128), connection: conn, handler: handler}, err
}

func (CR ConnReceiver) StartRecieving() {
	go func() {
		defer CR.connection.Close()
		desc := netpoll.Must(netpoll.HandleReadWrite(CR.connection))
		var err error
		for {
			CR.netPool.Start(desc, func(ev netpoll.Event) {

				if ev&netpoll.EventRead == 0 {
					CR.workPool.Schedule(func() {
						CR.handler.WriteNewMessage(CR.connection)
					})
				}

				CR.workPool.Schedule(func() {
					CR.handler.GetNewMessages(CR.connection)
				})

				if ev&netpoll.EventPollerClosed == 0 {
					err = errors.New("connection closed")
					return
				}

				fmt.Println("current event is : ", ev.String())
			})

			time.Sleep(10 * time.Millisecond)

			if err != nil {
				return
			}
		}
	}()
}
