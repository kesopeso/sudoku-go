package game

type CrossSolver struct {
	state *State
}

func NewCrossSolver(state *State) CrossSolver {
	return CrossSolver{
		state: state,
	}
}

func (s CrossSolver) GetSolutions(cell Position) []int {
	cellSolutions := s.state.GetCell(cell)
	solvedValues, unsolvedValues := s.getValues(cell)

	newCellSolutions := make([]int, 0, len(cellSolutions))
	newCellSolutionsToKeep := make([]int, 0, len(cellSolutions))

	for _, v := range cellSolutions {
		if solvedValues[v] {
			continue
		}

		newCellSolutions = append(newCellSolutions, v)

		if !unsolvedValues[v] {
			newCellSolutionsToKeep = append(newCellSolutionsToKeep, v)
		}
	}

	if len(newCellSolutionsToKeep) == 0 {
		return newCellSolutions
	}
	return newCellSolutionsToKeep
}

func (s CrossSolver) getValues(excludeCell Position) (solvedValues map[int]bool, unsolvedValues map[int]bool) {
	solvedValues = make(map[int]bool, s.state.size)
	unsolvedValues = make(map[int]bool, s.state.size)

	for i := range s.state.size {
		columnCellSolutions := s.state.GetCell(NewPosition(i, excludeCell.Column))
		rowCellSolutions := s.state.GetCell(NewPosition(excludeCell.Row, i))

		if i != excludeCell.Row {
			fillValues(columnCellSolutions, solvedValues, unsolvedValues)
		}

		if i != excludeCell.Column {
			fillValues(rowCellSolutions, solvedValues, unsolvedValues)
		}
	}

	for k := range solvedValues {
		delete(unsolvedValues, k)
	}

	return
}

func fillValues(solutions []int, solvedValues map[int]bool, unsolvedValues map[int]bool) {
	isSolved := len(solutions) == 1
	if isSolved {
		solvedValues[solutions[0]] = true
	} else {
		for _, ps := range solutions {
			unsolvedValues[ps] = true
		}
	}
}
