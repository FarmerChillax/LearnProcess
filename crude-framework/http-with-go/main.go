package main

import (
	"fin"
	"net/http"
)

func main() {
	app := fin.New()
	app.GET("/", func(ctx *fin.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Fin</h1>")
	})

	app.GET("/hello", func(ctx *fin.Context) {
		ctx.String(http.StatusOK, "hello, %s\n", ctx.Query("name"))
	})

	app.GET("/hello/:name", func(ctx *fin.Context) {
		ctx.String(http.StatusOK, "Hello %s, your're at %s\n", ctx.Param("name"), ctx.Path)
	})

	app.GET("/assets/*filepath", func(ctx *fin.Context) {
		ctx.JSON(http.StatusOK, fin.H{
			"filepath": ctx.Param("filepath"),
		})
	})

	app.Run("127.0.0.1:5000")
}
