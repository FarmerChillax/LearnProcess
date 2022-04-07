package main

import (
	"context"
	"grpc-demo/proto"
	"io"
	"log"
	"net"
	"runtime/debug"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamService struct{}

func (s *StreamService) List(r *proto.StreamRequest, stream proto.StreamService_ListServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.StreamResponse{
			Pt: &proto.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(i),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(stream proto.StreamService_RecordServer) error {
	defer log.Println("record stream was cloed!")
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("transport finish... waitting close")
				time.Sleep(time.Second * 3)
				return stream.SendAndClose(&proto.StreamResponse{
					Pt: &proto.StreamPoint{
						Name:  "close-name",
						Value: -1,
					},
				})
			}
			return err
		}
		log.Printf("server recv name: %s, value: %d \n", req.Pt.GetName(), req.Pt.GetValue())
	}
	return nil
}

func (s *StreamService) Route(stream proto.StreamService_RouteServer) error {
	return nil
}

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before gRPC method: %s, %v\n", info.FullMethod, req)
	// next
	resp, err := handler(ctx, req)

	log.Printf("after gRPC method %s, %v\n", info.FullMethod, resp)
	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", err)
		}
	}()

	return handler(ctx, req)
}

func main() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			RecoveryInterceptor,
			LoggingInterceptor,
		),
	}
	server := grpc.NewServer(opts...)
	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("net listen  err: %v\n", err)
	}

	server.Serve(lis)
}
