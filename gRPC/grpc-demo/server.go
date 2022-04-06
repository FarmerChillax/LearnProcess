package main

import (
	"context"
	"grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *proto.SearchRequest) (*proto.SearchResponse, error) {
	return &proto.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const ADDR = ":9001"

func main() {
	server := grpc.NewServer()
	proto.RegisterSearchServiceServer(server, &SearchService{})

	listen, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalf("net listen err: %v\n", err)
	}
	server.Serve(listen)
}
