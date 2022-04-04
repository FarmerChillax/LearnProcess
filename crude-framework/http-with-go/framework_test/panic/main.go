package main

import (
	"fin"
	"net/http"
)

func main() {
	app := fin.Default()
	app.GET("/", func(ctx *fin.Context) {
		ctx.String(http.StatusOK, "Hello Fin\n")
	})

	app.GET("/panic", func(ctx *fin.Context) {
		names := []string{"farmer"}
		ctx.String(http.StatusOK, names[123])
	})

	app.Run(":5000")
}
