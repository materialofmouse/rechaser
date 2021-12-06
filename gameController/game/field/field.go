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
	now_x = f.GetTargetX(x, direction, stride)
	now_y = f.GetTargetY(y, direction, stride)
	
	if (f.SetElement(GROUND, x, y) && f.SetElement(element, now_x, now_y)) {
		return f.GetElement(now_x, now_y)
	} else {
		return -1
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

func (f Field) GetTargetX(x int, direction int, length int) int {
	dx := []int{1, 0, -1, 0}
	x += dx[direction]
	
	return x
}

func (f Field) GetTargetY(y int, direction int, length int) int {
	dy := []int{0, 1, 0, -1}
	y += dy[direction]
	
	return y
}

func (f *Field) getValidPosition(x int, y int) bool {
	validX := (x >= 0 || x < len(f.getFields()))
	validY := (y >= 0 || y < len(f.getFields()[0]))
	
	return (validX && validY)
}

//要検討
func (f *Field) GetDirectionalArray(x int, y int, direction int, length int) []int {
	res := make([]int, length)

	for i := 0; i < length; i++ {
		x := f.GetTargetX(x, direction, i)
		y := f.GetTargetY(y, direction, i)
		element := f.GetElement(x, y)

		res = append(res, element)
	}
	return res
}

func (f *Field) GetArroundArray(x int, y int, length int) []int {
	res := make([]int, length*length)
	
	for i := 0; i < length; i++ {
		element := f.GetElement(x+i-1, y+i-1)

		res = append(res,element)
	}

	return res
}
