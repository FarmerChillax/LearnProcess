package main

import (
	"context"
	"fmt"
	"grpc-demo/pkg/gtls"
	"grpc-demo/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

const PORT = 5000

func main() {
	tlsClient := gtls.Client{
		ServerName: "grpc-demo",
		CertFile:   "../../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	//
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v\n", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	// cancel()
	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()
	resp, err := client.Search(ctx, &proto.SearchRequest{
		Request: "gRPC-clien-request",
	})
	fmt.Println(resp, err)
	// cancel()
	// if err != nil {
	// 	statusErr, ok := status.FromError(err)
	// 	if ok {
	// 		if statusErr.Code() == codes.DeadlineExceeded {
	// 			log.Fatalln("client.Search err: deadline")
	// 		}
	// 	}
	// 	log.Fatalf("client.Search err: %v\n", err)
	// }

	// log.Printf("resp: %s\n", resp.GetResponse())
}
