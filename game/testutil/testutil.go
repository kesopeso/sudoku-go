package testutil

import (
	"github.com/kesopeso/sudoku-go/game"
	"testing"
)

func InitSudoku(t *testing.T, cells [][]int) *game.Sudoku {
	t.Helper()
	sudoku, err := game.NewSudoku(cells)
	if err != nil {
		t.Fatalf("error occured: %v", err)
	}
	return sudoku
}

