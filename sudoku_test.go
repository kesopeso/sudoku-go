package main_test

import (
	"sudoku"
	"testing"
)

func TestFieldValidation(t *testing.T) {
	t.Run("too many elements in row defined", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4, 5},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		}
		fieldValidationHelper(t, fields, "too many elements in a row, found: 5, max: 4")
	})

	t.Run("incorrect number of rows defined", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		}
		fieldValidationHelper(t, fields, "fields array should be power of two, for instance 4, 9, 16, etc")
	})
}

func fieldValidationHelper(t *testing.T, fields [][]int, expectedError string) {
	t.Helper()
	_, error := main.NewSudoku(fields)
	if error == nil || error.Error() != expectedError {
		t.Fatalf("no error or bad error returned: %v", error)
	}
}
