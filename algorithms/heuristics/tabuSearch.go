package heuristics

import (
	"ixior/VehicleRoutingProblem/algorithms"
	"ixior/VehicleRoutingProblem/algorithms/greedy"
	"ixior/VehicleRoutingProblem/config"
	"ixior/VehicleRoutingProblem/model"
)

func TabuSearchHeuristic(data *model.CaseDTO, configFile *config.Config) *model.Solution {

	bestSolution := greedy.GetInitialSolution(data)
	currentSolution := bestSolution
	tabuList := make([]*model.Solution, 0)

	for i := 0; i < configFile.TabuSearchMaxIterations; i++ {
		solutions := algorithms.GetNeighbourSolutions(data, currentSolution, configFile)

		for _, solution := range solutions {
			if !contain(tabuList, solution) && solution.GetCost(data) < currentSolution.GetCost(data) {
				currentSolution = solution
			}
		}

		if currentSolution.GetCost(data) < bestSolution.GetCost(data) {
			bestSolution = currentSolution
		}
		tabuList = append(tabuList, currentSolution)

		if len(tabuList) > configFile.TabuSearchLengthOfTabuList {
			tabuList = tabuList[1:]
		}
	}

	return bestSolution
}

func contain(solutions []*model.Solution, newSolution *model.Solution) bool {
	for _, solution := range solutions {
		if equal(solution.Routes, newSolution.Routes) {
			return true
		}
	}
	return false
}

func equal(routes [][]int, routes2 [][]int) bool {
	if len(routes) != len(routes2) {
		return false
	}

	for i := 0; i < len(routes); i++ {
		if len(routes[i]) != len(routes2[i]) {
			return false
		}

		for j := 0; j < len(routes[i]); j++ {
			if routes[i][j] != routes[i][j] {
				return false
			}
		}
	}
	return true
}
