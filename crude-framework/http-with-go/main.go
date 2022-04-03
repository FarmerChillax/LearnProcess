package main

import (
	"fin"
	"fmt"
	"net/http"
)

func main() {
	app := fin.New()
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	app.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	app.Run("127.0.0.1:5000")
}
