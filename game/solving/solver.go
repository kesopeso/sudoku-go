package solving

import (
	"github.com/kesopeso/sudoku-go/game"
)

type CellSolver interface {
	SolveCell(s *game.Sudoku, row int, column int) []int
}

type Solver struct {
	cellSolvers []CellSolver
}

func NewSolver() Solver {
	cellSolvers := []CellSolver{
		NewSquareSolver(),
		NewRowSolver(),
		NewColumnSolver(),
		NewIntersectionSolver(),
	}
	return Solver{cellSolvers: cellSolvers}
}

func (so *Solver) Solve(s *game.Sudoku) {}
