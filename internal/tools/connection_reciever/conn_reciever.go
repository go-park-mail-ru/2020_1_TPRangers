package connection_reciever

import (
	"fmt"
	"github.com/mailru/easygo/netpoll"
	"main/internal/tools"
	"main/internal/tools/gpool"
	"net"
	"sync/atomic"
	"time"
)

type ConnReceiver struct {
	workPool   tools.GoPoolInterface
	netPool    netpoll.Poller
	connection net.Conn
	handler    tools.EventerInterface
	closedStatus int32
}

func NewConnReciever(conn net.Conn, handler tools.EventerInterface) (ConnReceiver, error) {
	poller, err := netpoll.New(nil)

	return ConnReceiver{netPool: poller, workPool: gpool.New(128), connection: conn, handler: handler , closedStatus: 0}, err
}

func (CR ConnReceiver) StartRecieving() {
	go func() {
		desc := netpoll.Must(netpoll.HandleReadWrite(CR.connection))

		CR.netPool.Start(desc, func(ev netpoll.Event) {

			fmt.Println("current event is : ", ev.String())

			if ev&netpoll.EventReadHup != 0 {
				CR.connection.Close()
				atomic.SwapInt32(&CR.closedStatus,1)
				return
			}

			if ev&netpoll.EventRead != 0 {
				CR.workPool.Schedule(func() {
					CR.handler.WriteNewMessage(CR.connection)
				})
			}


		})

		go func() {

			for {

				if val := atomic.SwapInt32(&CR.closedStatus,0) ; val == 1 {
					return
				}
				CR.workPool.Schedule(func() {
					CR.handler.GetNewMessages(CR.connection)
				})

				time.Sleep(10 * time.Second)
			}

		}()

	}()
}
