package main

import (
	"context"
	"grpc-demo/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9002", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("grpc dial err: %v\n", err)
	}
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)
	err = printLists(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "grpc stream client--> List",
			Value: 2022,
		},
	})
	if err != nil {
		log.Fatalf("printLists err: %v\n", err)
	}

	err = printRecord(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "grpc stream client--> record ",
			Value: 2022,
		},
	})
	if err != nil {
		log.Fatalf("printRecord err: %v\n", err)
	}

	err = printRoute(client, &proto.StreamRequest{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client--> Route",
			Value: 2022,
		},
	})
}

func printLists(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: name: %s, value: %d\n", resp.Pt.GetName(), resp.Pt.GetValue())
	}
	return nil
}

func printRecord(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 6; i++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
		r.Pt.Value++
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp.Name: %s, resp.Value: %d \n", resp.Pt.GetName(), resp.Pt.GetValue())
	return nil
}

func printRoute(client proto.StreamServiceClient, r *proto.StreamRequest) error {
	return nil
}
