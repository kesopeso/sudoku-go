package game_test

import (
	"github.com/kesopeso/sudoku-go/game"
	"reflect"
	"testing"
)

func TestCurrentSolutionRetrieval(t *testing.T) {
	t.Run("return current solution", func(t *testing.T) {
		sudoku := initSudoku(t, [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{2, 3, 0, 0},
			{0, 0, 0, 0},
		})

		got := sudoku.GetCurrentSolution()
		want := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{2, 3, 0, 0},
			{0, 0, 0, 0},
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("arrays are not the same, got: %v, want: %v", got, want)
		}
	})
}

func TestCellValidation(t *testing.T) {
	t.Run("too many elements in row defined", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 3, 4, 5},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		cellValidationErrorHelper(t, cells, "too few/many elements in a row, found: 5, expected: 4")
	})

	t.Run("incorrect number of rows defined", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		cellValidationErrorHelper(t, cells, "cells array should be power of two, for instance 4, 9, 16, etc")
	})

	t.Run("repeating numbers found in a row", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 4, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		cellValidationErrorHelper(t, cells, "row has repeating numbers")
	})

	t.Run("repeating numbers found in a column", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 3, 0},
		}
		cellValidationErrorHelper(t, cells, "column has repeating numbers")
	})

	t.Run("no repeating numbers in a row, ignore zeros", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{2, 3, 0, 0},
			{0, 0, 0, 0},
		}
		initSudoku(t, cells)
	})

	t.Run("no repeating numbers in a column, ignore zeros", func(t *testing.T) {
		cells := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{3, 0, 0, 0},
			{0, 0, 0, 0},
		}
		initSudoku(t, cells)
	})
}

func TestSolving(t *testing.T) {
	t.Run("return history with single (intial) entry when no possible solutions exist", func(t *testing.T) {
		cells := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		sudoku := initSudoku(t, cells)
		solutionHistory := sudoku.Solve()

		if len(solutionHistory) != 1 {
			t.Fatal("solution history does not have a single entry")
		}

		if len(solutionHistory[0].Changes) != 0 {
			t.Fatal("changes should be empty")
		}
	})
}

func cellValidationErrorHelper(t *testing.T, cells [][]int, expectedError string) {
	t.Helper()
	_, error := game.NewSudoku(cells)
	if error == nil || error.Error() != expectedError {
		t.Fatalf("no error or bad error returned: %v", error)
	}
}

func initSudoku(t *testing.T, cells [][]int) *game.Sudoku {
	t.Helper()
	sudoku, err := game.NewSudoku(cells)
	if err != nil {
		t.Fatalf("error occured: %v", err)
	}
	return sudoku
}
