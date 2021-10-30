package model

type Solution struct {
	Routes [][]int
	cost   float64
}

func (solution *Solution) AddRoute(route []int) {
	if solution.Routes == nil {
		solution.Routes = make([][]int, 1)
		solution.Routes[0] = route
	} else {
		solution.Routes = append(solution.Routes, route)
	}
	solution.cost = 0
}

func (solution *Solution) getRoute(i int) []int {
	if i > len(solution.Routes)-1 {
		return nil
	}
	return solution.Routes[i]
}

func (solution *Solution) GetCost(data *CaseDTO) float64 {
	if solution.cost == 0 {
		for i := 0; i < len(solution.Routes); i++ {
			solution.cost += data.Cost[0][solution.Routes[i][0]]
			for j := 0; j < len(solution.Routes[i])-1; j++ {
				solution.cost += data.Cost[solution.Routes[i][j]][solution.Routes[i][j+1]]
			}
			solution.cost += data.Cost[0][solution.Routes[i][len(solution.Routes[i])-1]]
		}
	}
	return solution.cost
}
