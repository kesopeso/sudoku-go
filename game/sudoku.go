package game

import (
	"fmt"
	"math"
)

type Sudoku struct {
	state *State
}

func NewSudoku(cells [][]int) (*Sudoku, error) {
	if error := validateInitalCells(cells); error != nil {
		return nil, error
	}
	state := initState(cells)
	return &Sudoku{state: state}, nil
}

func (s *Sudoku) GetSize() int {
	return s.state.size
}

func (s *Sudoku) GetSquareSize() int {
	sudokuSize := s.GetSize()
	squareSize := math.Sqrt(float64(sudokuSize))
	return int(squareSize)
}

func (s *Sudoku) GetCurrentSolution() [][]int {
	currentSolution := make([][]int, s.state.size)

	for r := range s.state.size {
		currentSolution[r] = make([]int, s.state.size)
		for c := range s.state.size {
			cellSolutions := s.state.GetCell(NewPosition(r, c))
			if len(cellSolutions) == 1 {
				currentSolution[r][c] = cellSolutions[0]
			}
		}
	}

	return currentSolution
}

func (s *Sudoku) GetCellState(row int, column int) (potentialSolutions []int, solution int, isSolved bool) {
	potentialSolutions = s.state.GetCell(NewPosition(row, column))
	if len(potentialSolutions) == 1 {
		solution = potentialSolutions[0]
	}
	isSolved = solution != 0
	return
}

func initState(cells [][]int) *State {
	tableSize := len(cells)
	stateTable := make([][][]int, tableSize)

	defaultCellState := make([]int, tableSize)
	for i := range tableSize {
		defaultCellState[i] = i + 1
	}

	for i := range tableSize {
		stateTable[i] = make([][]int, tableSize)
		for j := range tableSize {
			if len(cells[i]) > j {
				if cells[i][j] != 0 {
					stateTable[i][j] = []int{cells[i][j]}
					continue
				}
			}
			stateTable[i][j] = make([]int, tableSize)
			copy(stateTable[i][j], defaultCellState)
		}
	}

	return NewState(stateTable)
}

func validateInitalCells(cells [][]int) error {
	// check if cells size is 4, 9, 16, 25, 36, 49,...
	tableSize := len(cells)
	squareSize := math.Sqrt(float64(tableSize))
	if _, squareSizeDecimals := math.Modf(squareSize); squareSizeDecimals != 0 {
		return fmt.Errorf("cells array should be power of two, for instance 4, 9, 16, etc")
	}

	// validate repeating column numbers
	columnNumbers := make(map[int]map[int]bool, tableSize)
	var rowNumbers map[int]bool

	for _, rV := range cells {
		// validate repeating row numbers
		rowNumbers = make(map[int]bool, tableSize)

		// validate row length
		if len(rV) != tableSize {
			return fmt.Errorf("too few/many elements in a row, found: %v, expected: %v", len(rV), tableSize)
		}
		for cI, cV := range rV {
			// check if row number exists
			if rowNumbers[cV] {
				return fmt.Errorf("row has repeating numbers")
			}
			if cV != 0 {
				rowNumbers[cV] = true
			}

			// check if column number exists
			if _, ok := columnNumbers[cI]; !ok {
				columnNumbers[cI] = make(map[int]bool, tableSize)
			}
			if columnNumbers[cI][cV] {
				return fmt.Errorf("column has repeating numbers")
			}
			if cV != 0 {
				columnNumbers[cI][cV] = true
			}

			// validate cell min value
			if cV < 0 {
				return fmt.Errorf("cell cannot be smaller than 0 (which represents empty cell)")
			}
			// validate cell max value
			if cV > tableSize {
				return fmt.Errorf("cell cannot be bigger than %v", tableSize)
			}
		}
	}

	return nil
}
