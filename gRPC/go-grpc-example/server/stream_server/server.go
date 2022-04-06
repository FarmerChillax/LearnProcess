package main

import "go-grpc-example/proto"

type StreamService struct{}

func main() {

}

func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	return nil
}
