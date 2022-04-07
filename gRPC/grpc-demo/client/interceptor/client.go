package main

import (
	"context"
	"grpc-demo/pkg/gtls"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	tlsClient := gtls.Client{
		ServerName: "grpc-demo",
		CaFile:     "../../conf/ca.pem",
		CertFile:   "../../conf/client/client.pem",
		KeyFile:    "../../conf/client/client.key",
	}

	c, err := tlsClient.GetCredentialsByCA()
	if err != nil {
		log.Fatalf("GetTLSCredentialsByCA err: %v", err)
	}

	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(c))

	if err != nil {
		log.Fatalf("grpc.Dial err: %v\n", err)
	}

	defer conn.Close()

	// 传入 connect
	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatal("client search err: %v\n", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}
