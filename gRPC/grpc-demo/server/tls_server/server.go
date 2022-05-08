package main

import (
	"crypto/tls"
	"crypto/x509"
	"grpc-demo/proto"
	"io"
	"io/ioutil"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func main() {
	cert, err := tls.LoadX509KeyPair("../../conf/server/server.pem", "../../conf/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v\n", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v \n", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	server := grpc.NewServer(grpc.Creds(c))
	proto.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatalf("net listen  err: %v\n", err)
	}

	server.Serve(lis)

}
