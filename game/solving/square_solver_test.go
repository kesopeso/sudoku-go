package solving_test

import (
	"reflect"
	"sudoku/game/solving"
	"sudoku/game/testutil"
	"testing"
)

func TestSolveCellWithSquareSolver(t *testing.T) {
	t.Run("solve cell within square", func(t *testing.T) {
		sudoku := testutil.InitSudoku(t, [][]int{
			{0, 4, 3, 0, 0, 0, 0, 0, 8},
			{0, 0, 7, 8, 5, 1, 9, 0, 3},
			{0, 0, 8, 4, 0, 0, 0, 0, 2},
			{0, 7, 4, 6, 8, 0, 0, 0, 1},
			{0, 6, 1, 0, 3, 0, 8, 9, 0},
			{0, 3, 5, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 6, 0, 7, 2, 0},
			{0, 5, 2, 1, 9, 0, 3, 0, 0},
			{0, 8, 6, 0, 4, 0, 0, 1, 9},
		})

		squareSolver := solving.NewSquareSolver()
		got := squareSolver.SolveCell(sudoku, 5, 5)
		want := []int{2, 4, 5, 7, 9}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = squareSolver.SolveCell(sudoku, 7, 7)
		want = []int{4, 5, 6, 8}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}
	})
}
