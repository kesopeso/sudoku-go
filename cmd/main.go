package main

import (
	"fmt"
	"github.com/kesopeso/sudoku-go/game"
	"log"
)

func main() {
	var cells = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	sudoku, err := game.NewSudoku(cells)
	if err != nil {
		log.Fatal("error occured: ", err)
	}
	fmt.Println("this is sudoku setup", sudoku)
}
