package heuristics

import (
	"ixior462/VehicleRoutingProblem/algorithms"
	"ixior462/VehicleRoutingProblem/algorithms/greedy"
	"ixior462/VehicleRoutingProblem/config"
	"ixior462/VehicleRoutingProblem/model"
)

func LocalSearchHeuristic(data *model.CaseDTO, configFile *config.Config) *model.Solution {

	currentSolution := greedy.GetInitialSolution(data)
	improvement := 0

	for improvement < configFile.LocalSearchHeuristicImprovementCount {
		newSolution := algorithms.GetNeighbourSolution(data, currentSolution, configFile)

		if newSolution.GetCost(data) < currentSolution.GetCost(data) {
			improvement = 0
			currentSolution = newSolution
		} else {
			improvement++
		}
	}

	return currentSolution
}
