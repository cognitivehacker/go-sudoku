package grid

//New grid board
func New(sizeLine int, sizeColumn int) {

	board.sizeLine = sizeLine
	board.sizeColumn = sizeColumn
	board.grid = make(map[int]map[int]cell)
	// init grid
	for i := 0; i < board.sizeLine; i++ {
		board.grid[i] = make(map[int]cell)
		for j := 0; j < board.sizeColumn; j++ {
			board.grid[i][j] = cell{
				Value:         0,
				Possibilities: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			}
		}
	}

}
