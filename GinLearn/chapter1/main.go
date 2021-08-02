package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})
	// web
	r.GET("/user", queryString)
	// API demos
	r.GET("/user/:name", get)
	r.GET("/user/:name/*action", getRedirect)

	r.Run()
}

func queryString(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "Guest")
	lastName := c.Query("lastname") // 缩写 <- c.Request.URL.Query().Get("lastname")

	fmt.Println(c.QueryArray("lastname")) // 获取参数list
	if test, ok := c.GetQuery("lastname"); ok {
		fmt.Println(test)
	}
	lastName = c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

// API demo for get method.
func get(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(c.Params)
	c.JSON(200, gin.H{
		"msg":  fmt.Sprintf("Hello %s", name),
		"name": name,
	})
}

func getRedirect(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
