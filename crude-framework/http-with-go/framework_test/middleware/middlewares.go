package main

import (
	"fin"
	"log"
	"net/http"
	"time"
)

func Logger() fin.HandlerFunc {
	return func(ctx *fin.Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
func onlyForV2() fin.HandlerFunc {
	return func(ctx *fin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		ctx.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}
func main() {
	app := fin.New()
	app.Use(Logger())
	app.GET("/", func(ctx *fin.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Fin</h1>")
	})

	v2 := app.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/hello/:name", func(ctx *fin.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})

	app.Run(":5000")
}
