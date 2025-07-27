package solving_test

import (
	"github.com/kesopeso/sudoku-go/game/solving"
	"github.com/kesopeso/sudoku-go/game/testutil"
	"reflect"
	"testing"
)

func TestSolveCellWithRowSolver(t *testing.T) {
	t.Run("solve cell within row", func(t *testing.T) {
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

		rowSolver := solving.NewRowSolver()
		got := rowSolver.SolveCell(sudoku, 2, 4)
		want := []int{1, 3, 5, 6, 7, 9}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = rowSolver.SolveCell(sudoku, 7, 5)
		want = []int{4, 6, 7, 8}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}
	})
}

func TestSolveCellWithColumnSolver(t *testing.T) {
	t.Run("solve cell within column", func(t *testing.T) {
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

		columnSolver := solving.NewColumnSolver()
		got := columnSolver.SolveCell(sudoku, 1, 5)
		want := []int{1}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = columnSolver.SolveCell(sudoku, 5, 5)
		want = []int{2, 3, 4, 5, 6, 7, 8, 9}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}

		got = columnSolver.SolveCell(sudoku, 7, 8)
		want = []int{4, 5, 6, 7}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solutions don't match, got: %v, want: %v", got, want)
		}
	})
}
