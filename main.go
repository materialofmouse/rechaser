package main

import (
	//"net/http"
	//"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

	"rechaser/connect"
	"rechaser/game"
)

func main() {
	r := gin.Default()
	
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "home", })
	})
	var g game.Game

	//r.POST("/room", func())

	r.POST("/game", func(c *gin.Context) {
		var request connect.RequestData
		err := c.Bind(&request)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusBadRequest)
		}
		g.InitField(10)
		g.InitGame(request.Data.Session)
		g.InitPlayer("hoge", 0, 1, 1)
		fmt.Print("[LOG] ")
		fmt.Printf("SESSTION: %s | ", request.Data.Session)
		fmt.Printf("ACTION: %s | ", request.Data.Command.Action)
		fmt.Printf("DIRECTION: %s\n", request.Data.Command.Direction)
	})

	r.Run()
}
