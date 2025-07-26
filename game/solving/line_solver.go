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
		return solveRowCell(s, row, column)
	}
	return solveColumnCell(s, row, column)
}

func solveRowCell(s *game.Sudoku, row int, column int) []int {
	potentialSolutions, solution, isSolved := s.GetCellState(row, column)
	if isSolved {
		return []int{solution}
	}

	rowSolutions := getRowSolutions(s, row, column)

	solutions := make([]int, 0, len(potentialSolutions))
	for _, potentialSolution := range potentialSolutions {
		if rowSolutions[potentialSolution] {
			continue
		}
		solutions = append(solutions, potentialSolution)
	}

	return solutions
}

func solveColumnCell(s *game.Sudoku, row int, column int) []int {
	potentialSolutions, solution, isSolved := s.GetCellState(row, column)
	if isSolved {
		return []int{solution}
	}

	rowSolutions := getColumnSolutions(s, row, column)

	solutions := make([]int, 0, len(potentialSolutions))
	for _, potentialSolution := range potentialSolutions {
		if rowSolutions[potentialSolution] {
			continue
		}
		solutions = append(solutions, potentialSolution)
	}

	return solutions
}

func getRowSolutions(s *game.Sudoku, row int, column int) map[int]bool {
	sudokuSize := s.GetSize()
	rowSolutions := make(map[int]bool, sudokuSize)
	for c := range sudokuSize {
		if c == column {
			// this is the field itself, so it should not be checked
			continue
		}

		_, solution, isSolved := s.GetCellState(row, c)
		if !isSolved {
			continue
		}
		rowSolutions[solution] = true
	}
	return rowSolutions
}

func getColumnSolutions(s *game.Sudoku, row int, column int) map[int]bool {
	sudokuSize := s.GetSize()
	columnSolutions := make(map[int]bool, sudokuSize)
	for r := range sudokuSize {
		if r == row {
			// this is the field itself, so it should not be checked
			continue
		}

		_, solution, isSolved := s.GetCellState(r, column)
		if !isSolved {
			continue
		}
		columnSolutions[solution] = true
	}
	return columnSolutions
}
