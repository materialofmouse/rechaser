package gamecontroller

import (
	"rechaser/gameController/game"

	"github.com/google/uuid"
)

type gameController struct {
	rooms []*game.Game
}

func (g *gameController) getRoomCount() int {
	return len(g.rooms)
}

func (g *gameController) getRoom(i int) *game.Game {
	return g.rooms[i]
}

func (g *gameController) GetRoomFromID(_uuid uuid.UUID) (*game.Game, error){
	for i := 0; i < g.getRoomCount(); i++ {
		_g := g.getRoom(i)
		if _g.GetSessionID() == _uuid {
			return _g, nil 
		} 
	}
}


