package main

import (
	"context"
	"grpc-gw-demo/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("../certs/server.pem", "dev")
	if err != nil {
		log.Println("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(":50052", grpc.WithTransportCredentials(creds))
	defer conn.Close()
	if err != nil {
		log.Println(err)
	}

	c := proto.NewHelloWorldClient(conn)
	ctx := context.Background()
	body := &proto.HelloWorldRequest{
		Referer: "Grpc",
	}

	resp, err := c.SayHelloWorld(ctx, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(resp.Message)

}
