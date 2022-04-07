package main

import (
	"context"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())

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
