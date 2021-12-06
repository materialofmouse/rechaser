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

func (f *Field) SetElement(element int, x int , y int) bool {
	if (f.getValidPosition(x, y)) {
		f.fields[x][y] = element
		return true
	} else {
		return false
	}
}

func (f *Field) MoveElement(element int, x int, y int, direction int, stride int) int {
	now_x, now_y := x, y
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	
	now_x += dx[direction] * stride
	now_y += dy[direction] * stride
	
	if (f.SetElement(GROUND, x, y) && f.SetElement(element, now_x, now_y)) {
		return f.GetElement(now_x, now_y)
	} else {
		return -1
	}
}

func (f Field) collisionCheck(x int, y int, x1 int, y1 int) int {
	elm1 := f.GetElement(x, y)
	if (elm1 == WALL) {
		return WALL
	} else {	
		return GROUND
	}
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
	res := make([]int, length*length)
	
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			res = append(res, f.GetElement(x+i-1, y+j-1))
		}
	}

	return res
}
