package server

import (
	"context"
	"grpc-gw-demo/proto"
)

type HelloService struct{}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (h HelloService) SayHelloWorld(ctx context.Context, r *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{
		Message: r.GetReferer() + "server response",
	}, nil
}
