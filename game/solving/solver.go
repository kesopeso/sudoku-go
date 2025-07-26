package solving

import "sudoku/game"

// "math"

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

func (sudoku *Solver) Solve(s *game.Sudoku) {
	for {
		// run all cell solvers on each cell in parallel, get results from each, and use intersection to get possible results
		// if there are cells that have only one value, set that value and run again
		// if there are no cells with multiple values left, stop and return result
		// otherwise run in a decision tree and select the winning result
		// once is solved break
	}
}

// func (sudoku *Solver) getSolutionsIntersection(solutions [][]int) []int {
// 	solutionsCount := len(solutions)
// 	solutionsMap := make(map[int]int)
//
// 	for _, s := range solutions {
// 		for _, v := range s {
// 			solutionsMap[v]++
// 		}
// 	}
//
// 	solutionsIntersection := []int{}
// 	for k, v := range solutionsMap {
// 		if v == solutionsCount {
// 			solutionsIntersection = append(solutionsIntersection, k)
// 		}
// 	}
// 	return solutionsIntersection
// }
