package game

const (
	GROUND = iota
	WALL
	ENEMY
	ITEM
)

type field struct {
	fields [][]int
}

func (f *field) InitField(fieldRange int) {
	f.fields = make([][]int, fieldRange)
}

func (f field) GetFields() [][]int {
	return f.fields
}

func (f field) GetField(x int, y int) int {
	if len(f.fields) > x || len(f.fields) > y {
		return f.fields[x][y]
	} else {
		return -1
	}
}

