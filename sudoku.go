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

	for _, rV := range fields {
		// validate repeating row/column numbers

		// validate row length
		if len(rV) > fieldsCount {
			return fmt.Errorf("too many elements in a row, found: %v, max: %v", len(rV), fieldsCount)
		}
		for _, cV := range rV {
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
