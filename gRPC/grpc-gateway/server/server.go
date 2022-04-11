package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"grpc-gw-demo/pkg/ui/data/swagger"
	"grpc-gw-demo/pkg/utils"
	"grpc-gw-demo/proto"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	ServerPort  string
	CertName    string
	CertPemPath string
	CertKeyPath string
	EndPoint    string
	SwaggerDir  string
)

func Serve() (err error) {
	EndPoint = fmt.Sprintf(":%s", ServerPort)
	conn, err := net.Listen("tcp", EndPoint)
	fmt.Println("listen: ", conn.Addr().String())
	if err != nil {
		log.Printf("TCP Listen err: %v\n", err)
	}

	tlsConfig := utils.GetTLSConfig(CertPemPath, CertKeyPath)
	srv := createInternalServer(conn, tlsConfig)

	log.Printf("gRPC and https listen on: %s\n", ServerPort)

	if err = srv.Serve(tls.NewListener(conn, tlsConfig)); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	return err
}

func createInternalServer(conn net.Listener, tlsConfig *tls.Config) *http.Server {
	var opts []grpc.ServerOption

	// grpc server
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		log.Printf("Failed to create server TLS credentials %v", err)
	}

	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)

	// register grpc pb
	proto.RegisterHelloWorldServer(grpcServer, NewHelloService())

	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertName)
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()

	// register grpc-gateway pb

	if err := proto.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		log.Printf("Failed to register gw server: %v\n", err)
	}

	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   utils.GrpcHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	// fmt.Println(fileServer)
	prefix := "/swagger-ui"
	fmt.Println(prefix)
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s\n", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}
