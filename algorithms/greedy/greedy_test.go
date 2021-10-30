package greedy

import (
	"ixior462/VehicleRoutingProblem/model"
	"testing"
)

func TestInitialSolution(t *testing.T) {
	dto := model.CaseDTO{}
	dto.Cost = make([][]float64, 7)
	for i := 0; i < 7; i++ {
		dto.Cost[i] = make([]float64, 7)
	}
	dto.Capacity = 5

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			dto.Cost[i][j], dto.Cost[i][j] = 2, 2
		}
	}

	dto.Dimension = 7
	demands := make([]int, 0)
	demands = append(demands, 0, 2, 2, 2, 2, 2, 2)
	dto.Demands = demands

	initialSolution := GetInitialSolution(&dto)

	if initialSolution.Routes[0][0] != 1 && initialSolution.Routes[0][1] != 2 &&
		initialSolution.Routes[1][0] != 3 && initialSolution.Routes[1][1] != 4 &&
		initialSolution.Routes[2][0] != 5 && initialSolution.Routes[2][1] != 6 {
		t.Error()
	}
}
