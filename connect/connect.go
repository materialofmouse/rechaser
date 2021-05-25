package main

import (
	//"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "home", })
	})

	r.POST("/post", func(c *gin.Context) {
		fmt.Println(c.String)

	r.Run()
}
	
