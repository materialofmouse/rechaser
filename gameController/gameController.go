package gamecontroller

import (
	"fmt"
	"rechaser/gameController/game"

	"github.com/google/uuid"
)

type GameController struct{
	g []*game.Game
}

func (gc *GameController) CreateRoom() uuid.UUID {
	gameRoom := new(game.Game)
	_uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("please try again...")
	}
	gameRoom.InitGame(_uuid)
	gameRoom.InitField(10)
	gameRoom.SetPlayer("hoge", 0, 1, 1)
	gameRoom.SetPlayer("huga", 1, 2, 2)
	gc.g = append(gc.g, gameRoom)
	
	return _uuid
}

//後で最適化する
func (gc *GameController) GetRoom(sessionID uuid.UUID) *game.Game {
	for i := 0; i < len(gc.g); i++ {
		if gc.g[i].GetSessionID().String() == sessionID.String() {
			return gc.g[i]
		}
	}
	return nil
}
