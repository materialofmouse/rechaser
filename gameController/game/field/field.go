package field

const (
	GROUND = iota
	WALL
	ENEMY
	ITEM
)

type Field struct {
	fields [][]int
}

func (f *Field) InitField(fieldRange int) {
	f.fields = make([][]int, fieldRange)
}

func (f Field) getFields() [][]int {
	return f.fields
}

func (f Field) GetElement(x int, y int) int {
	if f.getValidPosition(x, y) {
		return f.fields[x][y]
	} else {
		return WALL
	}
}

func (f *Field) getValidPosition(x int, y int) bool {
	validX := (x >= 0 || x < len(f.getFields()))
	validY := (y >= 0 || y < len(f.getFields()[0]))
	
	return (validX && validY)
}

//要検討
func (f *Field) GetDirectionalArray(x int, y int, direction int, length int) []int {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	res := make([]int, length)
	for i := 0; i < length; i++ {
		x += dx[direction]
		y += dy[direction]

		res = append(res, f.GetElement(x, y))
	}

	return res
}

func (f *Field) GetArroundArray(x int, y int, length int) []int {
	res := make([]int, length)
	
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			res = append(res, f.GetElement(x+i-1, y+j-1))
		}
	}

	return res
}
