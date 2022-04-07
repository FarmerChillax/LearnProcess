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

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type SearchService struct {
	auth *Auth
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if err := s.auth.Check(ctx); err != nil {
		return nil, err
	}
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

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证失败")
	}

	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		log.Println("key:", value)
		appKey = value[0]
	}

	if value, ok := md["app_secret"]; ok {
		log.Println("secret:", value)
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return status.Errorf(codes.Unauthenticated, "自定义认证 token 无效")
	}

	return nil
}

func (a *Auth) GetAppKey() string {
	return "farmer"
}

func (a *Auth) GetAppSecret() string {
	return "12344321"
}
