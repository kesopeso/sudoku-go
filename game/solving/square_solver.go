package solving

import "sudoku/game"

type SquareSolver struct {
}

func NewSquareSolver() *SquareSolver {
	return &SquareSolver{}
}

func (ss *SquareSolver) SolveCell(s *game.Sudoku, row int, column int) []int {
	return nil
}
