package main

import "github.com/cognitivehacker/go-sudoku/pkg/grid"

func main() {

	grid.New(9, 9)
	grid.AddValue(0, 0, 999)
	grid.Print()
}
