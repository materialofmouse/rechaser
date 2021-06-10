package game

import (
	"fmt"
)

type Player struct{
	name string
	status int
	item int
	team int
}

func (p Player) GetName() string {
	return p.name
}

func (p *Player) SetName(_name string) {
	p.name = _name
}

func (p Player) GetItem() int {
	return p.item
}

func (p *Player) SetItem(item int) {
	p.item = item
}

func (p *Player) AddItem() {
	p.item += 1
}

func (p Player) GetStatus() int {
	return p.status
}

func (p *Player) SetStatus(status int) {
	p.status = status
}
