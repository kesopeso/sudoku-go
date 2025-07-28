package game

import (
	"fmt"
	"math"

	"github.com/kesopeso/sudoku-go/util"
)

type SquareSolver struct {
	state *State
}

type SquareBoundries struct {
	StartRow    int
	EndRow      int
	StartColumn int
	EndColumn   int
}

func NewSquareSolver(state *State) SquareSolver {
	return SquareSolver{
		state: state,
	}
}

func (s SquareSolver) GetSolutions(cell Position) []int {
	cellSolutions := s.state.GetCell(cell)
	solvedValues, unsolvedValues := s.getValues(cell)

	newCellSolutions := make([]int, 0, len(cellSolutions))
	newCellSolutionsToKeep := make([]int, 0, len(cellSolutions))

	for _, v := range cellSolutions {
		if solvedValues[v] {
			continue
		}

		newCellSolutions = append(newCellSolutions, v)

		if !unsolvedValues[v] {
			newCellSolutionsToKeep = append(newCellSolutionsToKeep, v)
		}
	}

	if len(newCellSolutionsToKeep) == 0 {
		return newCellSolutions
	}
	return newCellSolutionsToKeep
}

func (s SquareSolver) getValues(excludeCell Position) (solvedValues map[int]bool, unsolvedValues map[int]bool) {
	solvedValues = make(map[int]bool, s.state.size)
	unsolvedValues = make(map[int]bool, s.state.size)

	squareBoundries := s.getSquareBoundries(excludeCell)
	for r := squareBoundries.StartRow; r <= squareBoundries.EndRow; r++ {
		for c := squareBoundries.StartColumn; c <= squareBoundries.EndColumn; c++ {
			if r == excludeCell.Row && c == excludeCell.Column {
				continue
			}
			solutions := s.state.GetCell(NewPosition(r, c))
			isSolved := len(solutions) == 1
			if isSolved {
				solvedValues[solutions[0]] = true
			} else {
				for _, ps := range solutions {
					unsolvedValues[ps] = true
				}
			}
		}
	}

	for k := range solvedValues {
		delete(unsolvedValues, k)
	}

	return
}

func (s SquareSolver) getSquareBoundries(excludeCell Position) SquareBoundries {
	squareSize, remainder := math.Modf(math.Sqrt(float64(s.state.size)))
	util.Assert(remainder == 0, fmt.Sprintf("invalid sudoku size: %v", s.state.size))

	startRow := int(float64(excludeCell.Row) - math.Mod(float64(excludeCell.Row), squareSize))
	endRow := startRow + int(squareSize) - 1
	startColumn := int(float64(excludeCell.Column) - math.Mod(float64(excludeCell.Column), squareSize))
	endColumn := startColumn + int(squareSize) - 1

	return NewSquareBoundries(startRow, endRow, startColumn, endColumn)
}

func NewSquareBoundries(startRow int, endRow int, startColumn int, endColumn int) SquareBoundries {
	return SquareBoundries{
		StartRow:    startRow,
		EndRow:      endRow,
		StartColumn: startColumn,
		EndColumn:   endColumn,
	}
}
