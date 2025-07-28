package game

import (
	"fmt"

	"github.com/kesopeso/sudoku-go/util"
)

type State struct {
	size  int
	cells map[int]map[int][]int
}

type Position struct {
	Row    int
	Column int
}

func NewState(cells [][][]int) *State {
	stateCells := make(map[int]map[int][]int, len(cells))
	for r := range cells {
		for c := range cells[r] {
			if stateCells[r] == nil {
				stateCells[r] = make(map[int][]int, len(cells))
			}
			stateCells[r][c] = cells[r][c]
		}
	}

	return &State{
		size:  len(cells),
		cells: stateCells,
	}
}

func (s *State) GetCell(cell Position) []int {
	util.Assert(s.cells[cell.Row] != nil, fmt.Sprintf("cell position out of bounds, row: %v, column: %v", cell.Row, cell.Column))
	util.Assert(s.cells[cell.Row][cell.Column] != nil, fmt.Sprintf("cell position out of bounds, row: %v, column: %v", cell.Row, cell.Column))
	return s.cells[cell.Row][cell.Column]
}

func (s *State) SetCell(cell Position, solutions []int) {
	util.Assert(s.cells[cell.Row] != nil, fmt.Sprintf("cell position out of bounds, row: %v, column: %v", cell.Row, cell.Column))
	util.Assert(s.cells[cell.Row][cell.Column] != nil, fmt.Sprintf("cell position out of bounds, row: %v, column: %v", cell.Row, cell.Column))
	util.Assert(solutions != nil, "nil solutions parameter passed")
	s.cells[cell.Row][cell.Column] = solutions
}

func NewPosition(row int, column int) Position {
	return Position{
		Row:    row,
		Column: column,
	}
}

//
// func GetCrossValues(state *State, excludeCell Position) (solvedValues map[int]bool, unsolvedValues map[int]bool, err error) {
// 	solvedValues = make(map[int]bool, state.size)
// 	unsolvedValues = make(map[int]bool, state.size)
//
// 	for i := range state.size {
// 		columnCellSolutions, err := state.GetCell(NewPosition(i, excludeCell.Column))
// 		if err != nil {
// 			return nil, nil, err
// 		}
//
// 		rowCellSolutions, err := state.GetCell(NewPosition(excludeCell.Row, i))
// 		if err != nil {
// 			return nil, nil, err
// 		}
//
// 		if i != excludeCell.Row {
// 			fillValues(columnCellSolutions, solvedValues, unsolvedValues)
// 		}
//
// 		if i != excludeCell.Column {
// 			fillValues(rowCellSolutions, solvedValues, unsolvedValues)
// 		}
// 	}
// 	return
// }
//
// func fillValues(solutions []int, solvedValues map[int]bool, unsolvedValues map[int]bool) {
// 	isSolved := len(solutions) == 1
// 	if isSolved {
// 		solvedValues[solutions[0]] = true
// 	} else {
// 		for _, ps := range solutions {
// 			unsolvedValues[ps] = true
// 		}
// 	}
// }
