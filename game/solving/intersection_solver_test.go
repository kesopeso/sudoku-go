package solving_test

import (
	"reflect"
	"sudoku/game"
	"sudoku/game/solving"
	"testing"
)

func TestSolveCellWithIntersectionSolver(t *testing.T) {
	t.Run("solve cell with intersection", func(t *testing.T) {
		fields := [][]int{
			{0, 4, 3, 0, 0, 0, 0, 0, 8},
			{0, 0, 7, 8, 5, 1, 9, 0, 3},
			{0, 0, 8, 4, 0, 0, 0, 0, 2},
			{0, 7, 4, 6, 8, 0, 0, 0, 1},
			{0, 6, 1, 0, 3, 0, 8, 9, 0},
			{0, 3, 5, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 6, 0, 7, 2, 0},
			{0, 5, 2, 1, 9, 0, 3, 0, 0},
			{0, 8, 6, 0, 4, 0, 0, 1, 9},
		}
		sudoku, err := game.NewSudoku(fields)
		if err != nil {
			t.Fatalf("error occured: %v", err)
		}

		interactionSolver := solving.NewIntersectionSolver()
		got := interactionSolver.SolveCell(sudoku, 0, 4)
		want := []int{2, 7}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = interactionSolver.SolveCell(sudoku, 2, 4)
		want = []int{7}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}
	})
}
