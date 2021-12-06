package main

import (
	//"net/http"
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"rechaser/connect"
	"rechaser/gameController"
)

func main() {
	r := gin.Default()
	gc := new(gamecontroller.GameController)
	
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "home", })
	})

	r.GET("/create_room", func(c *gin.Context) {
		_uuid := gc.CreateRoom()
		c.JSON(http.StatusOK, gin.H{
			"message" : string("test"),
			"sessionID" : string(_uuid.String()),
		})
	})

	r.POST("/game", func(c *gin.Context) {
		var request connect.RequestData
		err := c.Bind(&request)
		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusBadRequest)
		}
		g := gc.GetRoom(request.Data.Session)
		if request.Data.Command.Action == "WALK" {
			g.GetReady()
		}
	})

	r.Run()
}
