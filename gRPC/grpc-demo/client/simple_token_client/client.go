package main

import (
	"context"
	"grpc-demo/pkg/gtls"
	"grpc-demo/proto"
	"log"

	"google.golang.org/grpc"
)

const PORT = 5000

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app_key":    a.AppKey,
		"app_secret": a.AppSecret,
	}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}

func main() {
	tlsClient := gtls.Client{
		ServerName: "grpc-demo",
		CertFile:   "../../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	auth := Auth{
		AppKey:    "farmer",
		AppSecret: "12344321",
	}
	//
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v\n", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchRequest{
		Request: "gRPC-clien-request",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v\n", err)
	}

	log.Printf("resp: %s\n", resp.GetResponse())
}
