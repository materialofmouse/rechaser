package game

import (
	"rechaser/game/player"
	"rechaser/game/field"
)

type Game struct{
	sessionID string
	f field.Field
	p []*player.Player
	turn int
	winner *player.Player
}

func (g *Game) InitGame(_sessionID string) {
	g.sessionID = _sessionID
}

func (g *Game) InitPlayer(_name string, _team int, _x int, _y int) bool {
	if len(g.p) <= 2{
		_p := player.NewPlayer()
		_p.InitPlayer(_name, _team, _x, _y)
		g.p = append(g.p, _p)
		return true
	} else {
		return false
	}
}

func (g *Game) InitField(fieldRange int) {
	g.f.InitField(fieldRange)
}

func (g *Game) Step() {
	g.turn++
	
}
