package solving

import "sudoku/game"

type LineSolver struct {
	options LineSolverOptions
}

type LineSolverOptions struct {
	solveRow bool
}

func newLineSolver(options LineSolverOptions) *LineSolver {
	return &LineSolver{options: options}
}

func NewRowSolver() *LineSolver {
	return newLineSolver(LineSolverOptions{solveRow: true})
}

func NewColumnSolver() *LineSolver {
	return newLineSolver(LineSolverOptions{solveRow: false})
}

func (ls *LineSolver) SolveCell(s *game.Sudoku, row int, column int) []int {
	if ls.options.solveRow {
		return ls.solveRowCell(s, row, column)
	}
	return ls.solveColumnCell(s, row, column)
}

func (ls *LineSolver) solveRowCell(s *game.Sudoku, row int, column int) []int {
	return nil
}

func (ls *LineSolver) solveColumnCell(s *game.Sudoku, row int, column int) []int {
	return nil
}
