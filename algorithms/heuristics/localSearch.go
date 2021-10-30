package heuristics

import (
	"ixior/VehicleRoutingProblem/algorithms"
	"ixior/VehicleRoutingProblem/algorithms/greedy"
	"ixior/VehicleRoutingProblem/config"
	"ixior/VehicleRoutingProblem/model"
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
