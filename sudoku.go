package main

import (
	"fmt"
	"math"
)

type Sudoku struct {
	stateTable [][][]int
}

func NewSudoku(fields [][]int) (*Sudoku, error) {
	if error := validateFields(fields); error != nil {
		return nil, error
	}
	stateTable := getStateTable(fields)
	return &Sudoku{stateTable: stateTable}, nil
}

func (s *Sudoku) GetCurrentSolution() [][]int {
	currentSolution := make([][]int, len(s.stateTable))

	for i, row := range s.stateTable {
		currentSolution[i] = make([]int, len(row))
		for j, v := range row {
			if len(v) == 1 {
				currentSolution[i][j] = v[0]
				continue
			}
		}
	}

	return currentSolution
}

func getStateTable(fields [][]int) [][][]int {
	tableSize := len(fields)
	stateTable := make([][][]int, tableSize)

	defaultCellState := make([]int, tableSize)
	for i := range tableSize {
		defaultCellState[i] = i + 1
	}

	for i := range tableSize {
		stateTable[i] = make([][]int, tableSize)
		for j := range tableSize {
			if len(fields[i]) > j {
				if fields[i][j] != 0 {
					stateTable[i][j] = []int{fields[i][j]}
					continue
				}
			}
			stateTable[i][j] = make([]int, tableSize)
			copy(stateTable[i][j], defaultCellState)
		}
	}

	return stateTable
}

func validateFields(fields [][]int) error {
	// check if field size 4, 9, 16, 25, 36, 49,...
	tableSize := len(fields)
	squareSize := math.Sqrt(float64(tableSize))
	if _, squareSizeDecimals := math.Modf(squareSize); squareSizeDecimals != 0 {
		return fmt.Errorf("fields array should be power of two, for instance 4, 9, 16, etc")
	}

	// validate repeating column numbers
	columnNumbers := map[int]map[int]bool{}
	var rowNumbers map[int]bool

	for _, rV := range fields {
		// validate repeating row numbers
		rowNumbers = map[int]bool{}

		// validate row length
		if len(rV) > tableSize {
			return fmt.Errorf("too many elements in a row, found: %v, max: %v", len(rV), tableSize)
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
				columnNumbers[cI] = map[int]bool{}
			}
			if columnNumbers[cI][cV] {
				return fmt.Errorf("column has repeating numbers")
			}
			if cV != 0 {
				columnNumbers[cI][cV] = true
			}

			// validate field min value
			if cV < 0 {
				return fmt.Errorf("field cannot be smaller than 0 (which represents empty field)")
			}
			// validate field max value
			if cV > tableSize {
				return fmt.Errorf("field cannot be bigger than %v", tableSize)
			}
		}
	}

	return nil
}
