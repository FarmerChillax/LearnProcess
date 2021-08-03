package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// set upload file
	router.Static("/", "./public")

	router.POST("/upload", formFile)

	router.POST("/fileBind", bindFile)

	router.Run()
}

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func bindFile(c *gin.Context) {
	// single file upload -- validator data
	var bindFile BindFile
	if err := c.ShouldBind(&bindFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	// save upload file
	file := bindFile.File
	dst := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"upload-err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name":  bindFile.Name,
		"Email": bindFile.Email,
		"file":  file.Filename,
		"dst":   dst,
	})
}

func formFile(c *gin.Context) {
	// single file upload
	// get file from FormFile
	file, err := c.FormFile("file")
	dst := filepath.Base(file.Filename)
	fmt.Println("base->", dst)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Upload file err: %s", err.Error()))
		return
	}
	fmt.Println(file.Filename, file.Size, file.Header)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errMsg": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"size":     file.Size,
		"header":   file.Header,
	})
}
