package solving

import (
	"github.com/kesopeso/sudoku-go/game"
	"math"
)

type SquareSolver struct {
}

func NewSquareSolver() *SquareSolver {
	return &SquareSolver{}
}

func (ss *SquareSolver) SolveCell(s *game.Sudoku, row int, column int) []int {
	potentialSolutions, solution, isSolved := s.GetCellState(row, column)
	if isSolved {
		return []int{solution}
	}

	startRow, endRow, startColumn, endColumn := getSquare(s, row, column)
	squareSolutions := getSquareSolutions(s, row, column, startRow, endRow, startColumn, endColumn)

	solutions := make([]int, 0, len(potentialSolutions))
	for _, potentialSolution := range potentialSolutions {
		if squareSolutions[potentialSolution] {
			continue
		}
		solutions = append(solutions, potentialSolution)
	}

	return solutions
}

func getSquareSolutions(s *game.Sudoku, row int, column int, startRow int, endRow int, startColumn int, endColumn int) map[int]bool {
	squareSolutions := make(map[int]bool, s.GetSize())
	for r := startRow; r <= endRow; r++ {
		for c := startColumn; c <= endColumn; c++ {
			// this is the field itself, so it should not be checked
			if r == row && c == column {
				continue
			}
			_, solution, isSolved := s.GetCellState(r, c)
			if !isSolved {
				continue
			}
			squareSolutions[solution] = true
		}
	}
	return squareSolutions
}

func getSquare(s *game.Sudoku, row int, column int) (startRow int, endRow int, startColumn int, endColumn int) {
	squareSize := s.GetSquareSize()
	startRow = int(float64(row) - math.Mod(float64(row), float64(squareSize)))
	endRow = startRow + squareSize - 1
	startColumn = int(float64(column) - math.Mod(float64(column), float64(squareSize)))
	endColumn = startColumn + squareSize - 1
	return
}
