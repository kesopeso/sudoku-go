package solving_test

import (
	"github.com/kesopeso/sudoku-go/game/solving"
	"github.com/kesopeso/sudoku-go/game/testutil"
	"reflect"
	"testing"
)

func TestSolveCellWithIntersectionSolver(t *testing.T) {
	t.Run("solve cell with intersection", func(t *testing.T) {
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

		interactionSolver := solving.NewIntersectionSolver()
		got := interactionSolver.SolveCell(sudoku, 7, 7)
		want := []int{8}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = interactionSolver.SolveCell(sudoku, 5, 0)
		want = []int{8}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}
	})
}
