package grid

type Board struct {
	sizeLine   int
	sizeColumn int
	grid       map[int]map[int]cell
}

type cell struct {
	Value         int
	Possibilities []int
}

var board = Board{}

func AddValue(line int, column int, number int) {
	cell := board.grid[line][column]
	cell.Value = number
	cell.Possibilities = []int{}
	board.grid[line][column] = cell
}
