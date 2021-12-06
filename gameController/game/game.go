package game

import (
	"rechaser/gameController/game/field"
	"rechaser/gameController/game/player"

	"github.com/google/uuid"
)

const (
	GETREADY = iota
	WALK
	PUT
	LOOK
	SEARCH
)

type Game struct{
	sessionID uuid.UUID
	f field.Field
	p []*player.Player
	turn int
	winner *player.Player
}

func (g *Game) InitGame(_sessionID uuid.UUID) {
	g.sessionID = _sessionID
}

func (g *Game) InitField(fieldRange int) {
	g.f.InitField(fieldRange)
}

func (g *Game) SetPlayer(_name string, _team int, _x int, _y int) bool {
	if len(g.p) <= 2{
		_p := player.NewPlayer()
		_p.InitPlayer(_name, _team, _x, _y)
		g.p = append(g.p, _p)
		return true
	} else {
		return false
	}
}

func (g *Game) GetPlayer(team int) *player.Player {
	if team == g.p[0].GetTeam() {
		return g.p[0]
	} else {
		return g.p[1]
	}
}

func (g *Game) GetSessionID() uuid.UUID {
	return g.sessionID
}

func (g *Game) TurnStep() {
	g.turn++
}

// actions
func (g *Game) GetReady(p *player.Player) []int {
	return g.f.GetArroundArray(p.GetX(), p.GetY(), 1)
}

func (g *Game) Walk(team int, direction int) []int {
	var p = g.GetPlayer(team)
	targetElement := g.f.MoveElement(field.ENEMY, p.GetX(), p.GetY(), direction, 1)
	p.Walk(direction, 1)
	
	if (targetElement == field.WALL) {
		p.SetStatus(0)
	} else if (targetElement == field.ITEM) {
		p.AddItem()
	}
	
	return g.f.GetArroundArray(p.GetX(), p.GetY(), 3)
}

func (g *Game) Put(team int, direction int) []int {
	var p = g.GetPlayer(team)
	x := p.GetX()
	y := p.GetY()
	g.f.SetElement(field.WALL, x, y)
	return g.f.GetArroundArray(p.GetX(), p.GetY(), 3)
}

func (g *Game) Look(team int, direction int) []int {
	var p = g.GetPlayer(team)
	return g.f.GetDirectionalArray(p.GetX(), p.GetY(), direction, 9)
}

func (g *Game) Search(team int, direction int) []int {
	var p = g.GetPlayer(team)
	x := g.f.GetTargetX(p.GetX(), direction, 1)
	y := g.f.GetTargetY(p.GetY(), direction, 1)
	return g.f.GetArroundArray(x, y, 3)
}

