package main

type CellSolver interface {
	SolveCell(s *Sudoku, x int, y int) []int
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

func (sudoku *Solver) Solve(s *Sudoku) {
	// while (solution not found)
	// run all cell solvers on each cell in parallel, get results from each, and use intersection to get possible results
	// if there are cells that have only one value, set that value and run again
	// if there are no cells with multiple values left, stop and return result
	// otherwise run in a decision tree and select the winning result
}

type SquareSolver struct {
}

func NewSquareSolver() *SquareSolver {
	return &SquareSolver{}
}

func (ss *SquareSolver) SolveCell(s *Sudoku, x int, y int) []int {
	return nil
}

type RowSolver struct {
}

func NewRowSolver() *RowSolver {
	return &RowSolver{}
}

func (rs *RowSolver) SolveCell(s *Sudoku, x int, y int) []int {
	return nil
}

type ColumnSolver struct {
}

func NewColumnSolver() *ColumnSolver {
	return &ColumnSolver{}
}

func (cs *ColumnSolver) SolveCell(s *Sudoku, x int, y int) []int {
	return nil
}

type IntersectionSolver struct {
}

func NewIntersectionSolver() *IntersectionSolver {
	return &IntersectionSolver{}
}

func (is IntersectionSolver) SolveCell(s *Sudoku, x int, y int) []int {
	return nil
}
