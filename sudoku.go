package main

import (
	"fmt"
	"math"
)

type Sudoku struct {
	Fields [][]int
}

func NewSudoku(fields [][]int) (*Sudoku, error) {
	if error := validateFields(fields); error != nil {
		return nil, error
	}
	return &Sudoku{Fields: fields}, nil
}

func validateFields(fields [][]int) error {
	// validate length of fields
	fieldsCount := len(fields)
	squareCount := math.Sqrt(float64(fieldsCount))
	squareCountAsInt := int(squareCount)
	if squareCount-float64(squareCountAsInt) != 0 {
		return fmt.Errorf("fields array should be power of two, for instance 4, 9, 16, etc")
	}

	// validate repeating column numbers
	columnNumbers := map[int]map[int]bool{}
	for _, rV := range fields {
		// validate repeating row numbers
		rowNumbers := map[int]bool{}

		// validate row length
		if len(rV) > fieldsCount {
			return fmt.Errorf("too many elements in a row, found: %v, max: %v", len(rV), fieldsCount)
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
			if cV > fieldsCount {
				return fmt.Errorf("field cannot be bigger than %v", fieldsCount)
			}
		}
	}

	return nil
}
