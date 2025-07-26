package main

import (
	"fmt"
	"log"
	"sudoku/game"
)

func main() {
	var fields = [][]int{
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
		{5, 3, 2, 1, 2, 2, 2, 2, 2},
	}
	sudoku, err := game.NewSudoku(fields)
	if err != nil {
		log.Fatal("error occured: ", err)
	}
	fmt.Println("this is sudoku setup", sudoku)
}
