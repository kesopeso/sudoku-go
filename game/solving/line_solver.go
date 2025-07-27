package solving

import "github.com/kesopeso/sudoku-go/game"

type LineSolver struct {
	options LineSolverOptions
}

type LineSolverOptions struct {
	solveRow bool
}

type CellStateGetter func(offset int) (solution int, isSolved bool)

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
	potentialSolutions, solution, isSolved := s.GetCellState(row, column)
	if isSolved {
		return []int{solution}
	}

	cellOffset, cellStateGetter := getCellOffsetAndStateGetter(s, row, column, ls.options.solveRow)
	lineSolutions := getLineSolutions(s, cellOffset, cellStateGetter)

	solutions := make([]int, 0, len(potentialSolutions))
	for _, potentialSolution := range potentialSolutions {
		if lineSolutions[potentialSolution] {
			continue
		}
		solutions = append(solutions, potentialSolution)
	}

	return solutions
}

func getCellOffsetAndStateGetter(s *game.Sudoku, row int, column int, isRowSolver bool) (int, CellStateGetter) {
	cell := row
	cellStateGetter := func(offset int) (solution int, isSolved bool) {
		_, solution, isSolved = s.GetCellState(offset, column)
		return
	}

	if isRowSolver {
		cell = column
		cellStateGetter = func(offset int) (solution int, isSolved bool) {
			_, solution, isSolved = s.GetCellState(row, offset)
			return
		}
	}

	return cell, cellStateGetter
}

func getLineSolutions(s *game.Sudoku, cellOffset int, getCellState CellStateGetter) map[int]bool {
	sudokuSize := s.GetSize()
	lineSolutions := make(map[int]bool, sudokuSize)
	for c := range sudokuSize {
		if c == cellOffset {
			// this is the field itself, so it should not be checked
			continue
		}

		solution, isSolved := getCellState(c)
		if !isSolved {
			continue
		}
		lineSolutions[solution] = true
	}
	return lineSolutions
}
