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
	getReadied bool
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) InitPlayer(_name string, _team int, _x int, _y int) {
	p.name = _name
	p.item = 0
	p.status = ALIVE
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

func (p *Player) GetX() int {
	return p.x
}

func (p *Player) GetY() int {
	return p.y
}

func (p *Player) Walk(direction int, stride int) {
	if direction == UP {
		p.y += stride
	} else if direction == DOWN {
		p.y -= stride
	} else if direction == RIGHT {
		p.x += stride
	} else if direction == LEFT {
		p.x -= stride
	}
}
