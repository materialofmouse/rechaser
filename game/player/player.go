package player

const (
	UP int = iota
	DOWN
	RIGHT
	LEFT
)

const (
	ALIVE = iota
	DIE
)

type Player struct{
	name string
	status int
	item int
	team int
	x int
	y int
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) InitPlayer(_name string, _team int, _x int, _y int) {
	p.name = _name
	p.item = 0
	p.team = _team
	p.x = _x
	p.y = _y
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

func (p Player) GetTeam() int {
	return p.team
}

func (p *Player) SetTeam(team int) {
	p.team = team
}

func (p *Player) Walk(direction int) {
	if direction == UP {
		p.y += 1
	} else if direction == DOWN {
		p.y -= 1
	} else if direction == RIGHT {
		p.x += 1
	} else if direction == LEFT {
		p.x -= 1
	}
}
