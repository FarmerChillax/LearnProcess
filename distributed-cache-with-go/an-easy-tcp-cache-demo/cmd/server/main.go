package main

import (
	"context"
	"fmt"
	"log"
	"tcp-cache-service/cache"
	"tcp-cache-service/tcp"
)

func startService(ctx context.Context, cacheServer *tcp.TCPServer) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	address := "http://127.0.0.1:2333"
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
	cacheServer := tcp.New(c)

	ctx := startService(context.Background(), cacheServer)
	<-ctx.Done()
	fmt.Println("Shutting down cache service.")
}
