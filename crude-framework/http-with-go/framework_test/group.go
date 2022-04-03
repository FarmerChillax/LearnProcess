package main

import (
	"fin"
	"net/http"
)

func main() {
	app := fin.New()

	app.GET("/index", func(ctx *fin.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := app.Group("/v1")
	{
		v1.GET("/", func(ctx *fin.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Fin</h1>")
		})

		v1.GET("/hello", func(ctx *fin.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := app.Group("/v2")
	v2.GET("/hello/:name", func(ctx *fin.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})
	v2.POST("/login", func(ctx *fin.Context) {
		ctx.JSON(http.StatusOK, fin.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	app.Run(":5000")
}
