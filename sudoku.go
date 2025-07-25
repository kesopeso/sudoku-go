package main

import "fmt"

type Sudoku struct {
	Count int
}

func NewSudoku() *Sudoku {
	return &Sudoku{
		Count: 10,
	}
}

func main() {
	sudoku := NewSudoku()
	fmt.Println("testing", sudoku.Count)
}
