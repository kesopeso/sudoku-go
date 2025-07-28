package game_test

import (
	"reflect"
	"testing"

	"github.com/kesopeso/sudoku-go/game"
)

func TestGetState(t *testing.T) {
	t.Run("get state values", func(t *testing.T) {
		state := initState(t)

		want := []int{1, 2}
		got := state.GetCell(game.NewPosition(0, 0))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("get cell value failed, want: %v, got: %v", want, got)
		}

		want = []int{1}
		got = state.GetCell(game.NewPosition(3, 3))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("get cell value failed, want: %v, got: %v", want, got)
		}
	})
}

func TestSetState(t *testing.T) {
	t.Run("set state values", func(t *testing.T) {
		state := initState(t)

		want := []int{1, 2, 3}

		position := game.NewPosition(2, 2)
		state.SetCell(position, want)

		got := state.GetCell(position)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("set cell value failed, want: %v, got: %v", want, got)
		}
	})
}

func initState(t *testing.T) *game.State {
	t.Helper()

	cells := [][][]int{
		{
			{1, 2},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		},
		{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		},
		{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
		},
		{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1},
		},
	}

	return game.NewState(cells)
}
