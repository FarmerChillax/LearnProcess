package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/", func(c *gin.Context) {
		// urlencode
		message := c.PostForm("msg")
		fmt.Println(message)
		c.JSON(http.StatusOK, gin.H{
			"msg": message,
		})
	})
	router.POST("/multipart", multipartHandler)
	router.Run()
}

func multipartHandler(c *gin.Context) {
	// postData := c.DefaultPostForm("name", "img")
	postType := c.DefaultPostForm("type", "jpg")
	fmt.Println(postType)
	c.JSON(http.StatusOK, gin.H{
		"type": postType,
	})
}
