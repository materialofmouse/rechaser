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

func (f Field) GetFields() [][]int {
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
	//後で書く
	return true
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
