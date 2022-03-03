package main

import (
	"context"
	"fmt"
	"http-cache-service-demo/cache"
	"http-cache-service-demo/server"
	"log"
)

func startService(ctx context.Context, cacheServer *server.Server) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	address := "http://127.0.0.1:233"
	go func() {
		log.Println()
		cacheServer.Listen()
		cancel()
	}()

	go func() {
		fmt.Printf("Cache service is running in %v\n", address)
		fmt.Println("Cache Service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		cancel()
	}()
	return ctx
}

func main() {
	c := cache.New()
	cacheServer := server.New(c)
	ctx := startService(context.Background(), cacheServer)
	<-ctx.Done()
	fmt.Println("Shutting down cache service.")
}
