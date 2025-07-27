package solving

import (
	"github.com/kesopeso/sudoku-go/game"
)

type IntersectionSolver struct {
}

type CellPosition struct {
	Row    int
	Column int
}

func NewIntersectionSolver() *IntersectionSolver {
	return &IntersectionSolver{}
}

func (is IntersectionSolver) SolveCell(s *game.Sudoku, row int, column int) []int {
	potentialSolutions, solution, isSolved := s.GetCellState(row, column)
	if isSolved {
		return []int{solution}
	}

	// this removes values that are in the same row and column
	filteredPotentialSolutions := getFilteredPotentialSolutions(s, row, column, potentialSolutions)
	// this gets all the cells in the same square that are not solved
	unsolvedSquareCells := getUnsolvedSquareCells(s, row, column)

	intersectionsFound := make(map[int]int, len(filteredPotentialSolutions))
	for _, unsolvedSquareCell := range unsolvedSquareCells {
		intersectionSolutions := getIntersectionSolutions(s, unsolvedSquareCell.Row, unsolvedSquareCell.Column)
		for _, potentialSolution := range filteredPotentialSolutions {
			if intersectionSolutions[potentialSolution] {
				intersectionsFound[potentialSolution]++
			}
			if intersectionsFound[potentialSolution] == len(unsolvedSquareCells) {
				return []int{potentialSolution}
			}
		}
	}

	return filteredPotentialSolutions
}

func getFilteredPotentialSolutions(s *game.Sudoku, row int, column int, potentialSolutions []int) []int {
	intersectionSolutions := getIntersectionSolutions(s, row, column)
	filteredPotentialSolutions := make([]int, 0, len(potentialSolutions))
	for _, s := range potentialSolutions {
		if intersectionSolutions[s] {
			continue
		}
		filteredPotentialSolutions = append(filteredPotentialSolutions, s)
	}
	return filteredPotentialSolutions
}

func getIntersectionSolutions(s *game.Sudoku, row int, column int) map[int]bool {
	sudokuSize := s.GetSize()
	intersectionSolutions := make(map[int]bool, sudokuSize)
	for i := range sudokuSize {
		if i != column {
			_, solution, isSolved := s.GetCellState(row, i)
			if isSolved {
				intersectionSolutions[solution] = true
			}
		}
		if i != row {
			_, solution, isSolved := s.GetCellState(i, column)
			if isSolved {
				intersectionSolutions[solution] = true
			}
		}
	}
	return intersectionSolutions
}

func getUnsolvedSquareCells(s *game.Sudoku, row int, column int) []CellPosition {
	unsolvedSquareCells := make([]CellPosition, s.GetSize())
	startRow, endRow, startColumn, endColumn := getSquare(s, row, column)
	for r := startRow; r <= endRow; r++ {
		for c := startColumn; c <= endColumn; c++ {
			if r == row && c == column {
				continue
			}
			_, _, isSolved := s.GetCellState(r, c)
			if isSolved {
				continue
			}
			unsolvedSquareCells = append(unsolvedSquareCells, CellPosition{Row: r, Column: c})
		}
	}
	return unsolvedSquareCells
}
