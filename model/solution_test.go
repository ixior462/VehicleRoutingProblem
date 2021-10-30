package model

import (
	"testing"
)

func TestGetCost(t *testing.T) {
	dto := CaseDTO{}

	dto.Cost = make([][]float64, 7)
	for i := 0; i < 7; i++ {
		dto.Cost[i] = make([]float64, 7)
	}

	dto.Cost[0][1], dto.Cost[1][0] = 2, 2
	dto.Cost[1][2], dto.Cost[2][1] = 3, 3
	dto.Cost[0][2], dto.Cost[2][0] = 5, 5

	dto.Cost[0][3], dto.Cost[3][0] = 7, 7

	dto.Cost[0][4], dto.Cost[4][0] = 11, 11
	dto.Cost[4][5], dto.Cost[5][4] = 13, 13
	dto.Cost[5][6], dto.Cost[6][5] = 17, 17
	dto.Cost[0][6], dto.Cost[6][0] = 19, 19

	solution := Solution{Routes: [][]int{{1, 2}, {3}, {4, 5, 6}}}

	cost := solution.GetCost(&dto)

	if cost != 2+3+5+7+7+11+13+17+19 {
		t.Error(cost, " != ", 2+3+5+7+7+11+13+17+19)
	}
}
func TestAddRoute(t *testing.T) {
	route1 := make([]int, 0)
	route1 = append(route1, 1)
	route1 = append(route1, 2)

	route2 := make([]int, 0)
	route2 = append(route2, 3)
	route2 = append(route2, 4)

	solution := Solution{}
	solution.AddRoute(route1)
	solution.AddRoute(route2)

	if !equalSlice(route1, solution.getRoute(0)) {
		t.Error()
	}
	if !equalSlice(route2, solution.getRoute(1)) {
		t.Error()
	}

}

func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
