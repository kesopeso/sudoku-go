package game_test

import (
	"reflect"
	"sudoku/game"
	"sudoku/game/testutil"
	"testing"
)

func TestCurrentSolutionRetrieval(t *testing.T) {
	t.Run("return current solution", func(t *testing.T) {
		sudoku := testutil.InitSudoku(t, [][]int{
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

func TestFieldValidation(t *testing.T) {
	t.Run("too many elements in row defined", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4, 5},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		fieldValidationErrorHelper(t, fields, "too few/many elements in a row, found: 5, expected: 4")
	})

	t.Run("incorrect number of rows defined", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		fieldValidationErrorHelper(t, fields, "fields array should be power of two, for instance 4, 9, 16, etc")
	})

	t.Run("repeating numbers found in a row", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 4, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}
		fieldValidationErrorHelper(t, fields, "row has repeating numbers")
	})

	t.Run("repeating numbers found in a column", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 3, 0},
		}
		fieldValidationErrorHelper(t, fields, "column has repeating numbers")
	})

	t.Run("no repeating numbers in a row, ignore zeros", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{2, 3, 0, 0},
			{0, 0, 0, 0},
		}
		testutil.InitSudoku(t, fields)
	})

	t.Run("no repeating numbers in a column, ignore zeros", func(t *testing.T) {
		fields := [][]int{
			{1, 2, 3, 4},
			{0, 0, 0, 0},
			{3, 0, 0, 0},
			{0, 0, 0, 0},
		}
		testutil.InitSudoku(t, fields)
	})
}

func fieldValidationErrorHelper(t *testing.T, fields [][]int, expectedError string) {
	t.Helper()
	_, error := game.NewSudoku(fields)
	if error == nil || error.Error() != expectedError {
		t.Fatalf("no error or bad error returned: %v", error)
	}
}
