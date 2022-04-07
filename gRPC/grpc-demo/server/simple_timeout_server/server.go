package main

import (
	"context"
	"fmt"
	"grpc-demo/pkg/gtls"
	"grpc-demo/proto"
	pb "grpc-demo/proto"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	log.Println("start sleep...")
	time.Sleep(time.Second * 3)
	for i := 0; i < 5; i++ {
		log.Printf("sleep %d second. ctx: %v --> %v", i+1, ctx.Err(), ctx.Done())
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
		}
	}
	log.Println("I was wakeup !!!")
	return &pb.SearchResponse{Response: r.GetRequest() + " HTTP Server"}, nil
}

const PORT = 5000

func main() {
	certFile := "../../conf/server/server.pem"
	keyFile := "../../conf/server/server.key"
	tlsServer := gtls.Server{
		CertFile: certFile,
		KeyFile:  keyFile,
	}

	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	mux := GetHTTPServeMux()
	server := grpc.NewServer(grpc.Creds(c))

	proto.RegisterSearchServiceServer(server, &SearchService{})

	http.ListenAndServeTLS(fmt.Sprintf(":%d", PORT), certFile, keyFile,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				server.ServeHTTP(w, r)
			} else {
				mux.ServeHTTP(w, r)
			}
		}))

}

func GetHTTPServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("grpc-demo"))
	})
	return mux
}
