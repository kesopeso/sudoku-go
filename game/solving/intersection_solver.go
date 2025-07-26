package solving

import (
	"sudoku/game"
)

type IntersectionSolver struct {
}

func NewIntersectionSolver() *IntersectionSolver {
	return &IntersectionSolver{}
}

func (is IntersectionSolver) SolveCell(s *game.Sudoku, row int, column int) []int {
	return nil
}
