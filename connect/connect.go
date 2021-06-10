package main

import (
	//"net/http"
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
	Data struct {
		Session string `json:"session"`
		Command struct {
			Action string `json:"action"`
			Direction string `json:"direction"`
		} `json:"command"`
	} `json:"data"`
}

func main() {
	r := gin.Default()
	
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "home", })
	})

	r.POST("/post", func(c *gin.Context) {
		fmt.Println("post")
		var request RequestData
		err := c.Bind(&request)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusBadRequest)
		}
		fmt.Println(request.Data.Session)
		fmt.Println(request.Data.Command.Action)
	})

	r.Run()
}
	
