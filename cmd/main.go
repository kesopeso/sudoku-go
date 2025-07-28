package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kesopeso/sudoku-go/game"
)

func main() {
	var cells = [][]int{
		{6, 7, 0, 0, 0, 0, 0, 5, 1},
		{0, 0, 0, 8, 0, 0, 0, 0, 6},
		{0, 0, 0, 2, 0, 0, 4, 0, 0},
		{0, 9, 1, 0, 0, 0, 0, 0, 8},
		{7, 0, 0, 0, 0, 5, 0, 0, 0},
		{0, 0, 5, 6, 0, 3, 1, 0, 0},
		{0, 0, 0, 1, 0, 2, 0, 0, 0},
		{0, 1, 0, 0, 9, 0, 7, 3, 0},
		{0, 0, 3, 7, 0, 0, 0, 0, 0},
	}
	sudoku, err := game.NewSudoku(cells)
	if err != nil {
		log.Fatal("error occured: ", err)
	}

	solutionHistory := sudoku.Solve()
	for i, sh := range solutionHistory {
		fmt.Print("change " + strconv.Itoa(i))

		unsolvedCellsCount := 0
		for _, row := range sh.CurrentSolution {
			for _, cellSolution := range row {
				if cellSolution == 0 {
					unsolvedCellsCount++
				}
			}
		}
		if unsolvedCellsCount > 0 {
			fmt.Println(" (not solved)")
		} else {
			fmt.Println(" (solved)")
		}

		for _, c := range sh.Changes {
			if len(c.Solutions) > 1 {
				continue
			}
			fmt.Printf("row: %v, column: %v, new solution: %v\n", c.Cell.Row+1, c.Cell.Column+1, c.Solutions)
		}
	}
}
