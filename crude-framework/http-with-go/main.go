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

	app.POST("/login", func(ctx *fin.Context) {
		ctx.JSON(http.StatusOK, fin.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	app.Run("127.0.0.1:5000")
}
