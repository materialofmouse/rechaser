package main

import (
	//"net/http"
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"rechaser/connect"
	"rechaser/game"
)

func main() {
	r := gin.Default()
	
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "home", })
	})
	g := make([]*game.Game,0)
	
	//r.POST("/room", func())

	r.GET("/create_room", func(c *gin.Context) {
		var _g game.Game
		var message string
		
		_g.InitField(10)
		_uuid, err := uuid.NewRandom()
		fmt.Println(_uuid)
		if err != nil {
			message = "plase try again..."
		} else {
			message = ""
		}
		
		_g.InitGame(_uuid)
		g = append(g, &_g)
		
		c.JSON(http.StatusOK, gin.H{
			"message" : string(message),
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
	})

	r.Run()
}
