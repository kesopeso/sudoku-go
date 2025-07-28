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
