package solving_test

import (
	"testing"

	"github.com/kesopeso/sudoku-go/game/solving"
	"github.com/kesopeso/sudoku-go/game/testutil"
)

func TestSolveOneIteration(t *testing.T) {
	t.Run("solve cell which gets result from the row solver", func(t *testing.T) {
		sudoku := testutil.InitSudoku(t, [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 2, 3, 4, 0, 6, 9, 7, 8},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		})
		solver := solving.NewSolver()
		solver.Solve(sudoku)
	})

	t.Run("solve cell which gets result from the column solver", func(t *testing.T) {
	})
}
