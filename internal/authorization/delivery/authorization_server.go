package delivery
import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3080")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()


	fmt.Println("starting server at :8081")
	server.Serve(lis)
}