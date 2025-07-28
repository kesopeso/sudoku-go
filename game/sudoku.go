package game

import (
	"fmt"
	"math"

	"github.com/kesopeso/sudoku-go/util"
)

type Sudoku struct {
	state   *State
	solvers []Solver
}

type Solver interface {
	GetSolutions(cell Position) []int
}

type CellSolution struct {
	Cell      Position
	Solutions []int
}

type SolutionHistory struct {
	CurrentSolution [][]int
	Changes         []CellSolution
}

func NewCellSolution(cell Position, solutions []int) CellSolution {
	return CellSolution{
		Cell:      cell,
		Solutions: solutions,
	}
}

func NewSolutionHistory(currentSolution [][]int, changes []CellSolution) SolutionHistory {
	return SolutionHistory{
		CurrentSolution: currentSolution,
		Changes:         changes,
	}
}

func NewSudoku(cells [][]int) (*Sudoku, error) {
	if error := validateInitalCells(cells); error != nil {
		return nil, error
	}
	state := initState(cells)

	solvers := []Solver{
		NewCrossSolver(state),
		NewSquareSolver(state),
	}

	return &Sudoku{state: state, solvers: solvers}, nil
}

func (s *Sudoku) GetCurrentSolution() [][]int {
	currentSolution := make([][]int, s.state.size)

	for r := range s.state.size {
		currentSolution[r] = make([]int, s.state.size)
		for c := range s.state.size {
			cellSolutions := s.state.GetCell(NewPosition(r, c))
			if len(cellSolutions) == 1 {
				currentSolution[r][c] = cellSolutions[0]
			}
		}
	}

	return currentSolution
}

func (s *Sudoku) Solve() []SolutionHistory {
	solutionHistory := []SolutionHistory{
		NewSolutionHistory(s.GetCurrentSolution(), []CellSolution{}),
	}

	for {
		unsolvedCells := s.getUnsolvedCells()

		unsolvedCellsCount := len(unsolvedCells)
		if unsolvedCellsCount == 0 {
			fmt.Println("sudoku solved!")
			break
		}

		changesCh := make(chan CellSolution, unsolvedCellsCount)

		for _, uc := range unsolvedCells {
			go (func(unsolvedCell Position) {
				newSolutions := [][]int{}
				for _, solver := range s.solvers {
					newSolutions = append(newSolutions, solver.GetSolutions(unsolvedCell))
				}
				newSolutionsIntersection := util.ArraysIntersection(newSolutions...)
				changesCh <- NewCellSolution(unsolvedCell, newSolutionsIntersection)
			})(uc)
		}

		changes := []CellSolution{}
		for range unsolvedCellsCount {
			newCellSolution := <-changesCh
			currentCellSolution := s.state.GetCell(newCellSolution.Cell)
			util.Assert(
				len(newCellSolution.Solutions) <= len(currentCellSolution),
				fmt.Sprintf(
					"new cell solution has more entries than the current one, new: %v, current: %v",
					newCellSolution.Solutions,
					currentCellSolution,
				),
			)

			if len(newCellSolution.Solutions) < len(currentCellSolution) {
				changes = append(changes, newCellSolution)
			}
		}

		// no solutions found... we need to break
		if len(changes) == 0 {
			fmt.Println("no more solutions found, breaking out")
			break
		}

		for _, cs := range changes {
			s.state.SetCell(cs.Cell, cs.Solutions)
		}
		currentSolution := s.GetCurrentSolution()
		solutionHistory = append(solutionHistory, NewSolutionHistory(currentSolution, changes))
	}

	return solutionHistory
}

func (s *Sudoku) GetCellState(row int, column int) (potentialSolutions []int, solution int, isSolved bool) {
	potentialSolutions = s.state.GetCell(NewPosition(row, column))
	if len(potentialSolutions) == 1 {
		solution = potentialSolutions[0]
	}
	isSolved = solution != 0
	return
}

func (s *Sudoku) getUnsolvedCells() []Position {
	unsolvedCells := make([]Position, 0, s.state.size*s.state.size)
	for r := range s.state.size {
		for c := range s.state.size {
			cellPosition := NewPosition(r, c)
			cellSolution := s.state.GetCell(cellPosition)
			if len(cellSolution) > 1 {
				unsolvedCells = append(unsolvedCells, cellPosition)
			}
		}
	}
	return unsolvedCells
}

func initState(cells [][]int) *State {
	tableSize := len(cells)
	stateTable := make([][][]int, tableSize)

	defaultCellState := make([]int, tableSize)
	for i := range tableSize {
		defaultCellState[i] = i + 1
	}

	for i := range tableSize {
		stateTable[i] = make([][]int, tableSize)
		for j := range tableSize {
			if len(cells[i]) > j {
				if cells[i][j] != 0 {
					stateTable[i][j] = []int{cells[i][j]}
					continue
				}
			}
			stateTable[i][j] = make([]int, tableSize)
			copy(stateTable[i][j], defaultCellState)
		}
	}

	return NewState(stateTable)
}

func validateInitalCells(cells [][]int) error {
	// check if cells size is 4, 9, 16, 25, 36, 49,...
	tableSize := len(cells)
	squareSize := math.Sqrt(float64(tableSize))
	if _, squareSizeDecimals := math.Modf(squareSize); squareSizeDecimals != 0 {
		return fmt.Errorf("cells array should be power of two, for instance 4, 9, 16, etc")
	}

	// validate repeating column numbers
	columnNumbers := make(map[int]map[int]bool, tableSize)
	var rowNumbers map[int]bool

	for _, rV := range cells {
		// validate repeating row numbers
		rowNumbers = make(map[int]bool, tableSize)

		// validate row length
		if len(rV) != tableSize {
			return fmt.Errorf("too few/many elements in a row, found: %v, expected: %v", len(rV), tableSize)
		}
		for cI, cV := range rV {
			// check if row number exists
			if rowNumbers[cV] {
				return fmt.Errorf("row has repeating numbers")
			}
			if cV != 0 {
				rowNumbers[cV] = true
			}

			// check if column number exists
			if _, ok := columnNumbers[cI]; !ok {
				columnNumbers[cI] = make(map[int]bool, tableSize)
			}
			if columnNumbers[cI][cV] {
				return fmt.Errorf("column has repeating numbers")
			}
			if cV != 0 {
				columnNumbers[cI][cV] = true
			}

			// validate cell min value
			if cV < 0 {
				return fmt.Errorf("cell cannot be smaller than 0 (which represents empty cell)")
			}
			// validate cell max value
			if cV > tableSize {
				return fmt.Errorf("cell cannot be bigger than %v", tableSize)
			}
		}
	}

	return nil
}
