package game_test

import (
	"reflect"
	"testing"

	"github.com/kesopeso/sudoku-go/game"
)

func TestSquareSolver(t *testing.T) {
	t.Run("calculate new square solutions", func(t *testing.T) {
		cells := [][][]int{
			{
				{1},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			{
				{1, 2, 3, 4},
				{2},
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			},
			{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{3, 4},
				{3, 4},
			},
			{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
				{2, 3, 4},
				{1, 3, 4},
			},
		}

		state := game.NewState(cells)
		squareSolver := game.NewSquareSolver(state)

		want := []int{3, 4}
		got := squareSolver.GetSolutions(game.NewPosition(0, 1))

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("solution calculation failed, want: %v, got: %v", want, got)
		}

		want = []int{1}
		got = squareSolver.GetSolutions(game.NewPosition(3, 3))

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("solution calculation failed, want: %v, got: %v", want, got)
		}
	})
}
