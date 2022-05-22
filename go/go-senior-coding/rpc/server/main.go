package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
)

type HelloService struct{}

func (hs *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	go func() {
		log.Println("pprof star in :5122")
		http.ListenAndServe(":5122", nil)
	}()
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":9980")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		var input string
		log.Println("RPC Service running in :9980 ...")
		log.Println("Enter anything to cancel.")
		fmt.Scanln(&input)
		cancel()
	}()

	go StartService(ctx, listener)

	<-ctx.Done()

}

func StartService(ctx context.Context, listener net.Listener) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			go rpc.ServeConn(conn)
		}
	}
}
