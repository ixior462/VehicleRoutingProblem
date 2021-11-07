package algorithms

import (
	"ixior462/VehicleRoutingProblem/algorithms/transformations"
	"ixior462/VehicleRoutingProblem/config"
	"ixior462/VehicleRoutingProblem/model"
	"math/rand"
)

func GetNeighbourSolutions(data *model.CaseDTO, inputSolution *model.Solution, configFile *config.Config) []*model.Solution {

	var solutions = make([]*model.Solution, 0)

	if rand.Float64() < configFile.DeterministicNeighborTransformationOccurrenceChance {
		for i := 0; i < configFile.ShortRouteImprovementTransformationIterations; i++ {
			solutions = append(solutions, transformations.ShortRouteImprovementTransformation(data, inputSolution))
		}
		for i := 0; i < configFile.ReplaceHighestAverageNeighborhoodTransformationIterations; i++ {
			solutions = append(solutions, transformations.ReplaceHighestAverageNeighborhoodTransformation(data, inputSolution))
		}
		for i := 0; i < configFile.HighCostNodeImprovementTransformationIterations; i++ {
			solutions = append(solutions, transformations.HighCostNodeImprovementTransformation(data, inputSolution))
		}
	}

	for i := 0; i < configFile.RandomRouteImprovementTransformationIterations; i++ {
		solutions = append(solutions, transformations.RandomRouteImprovementTransformation(data, inputSolution, configFile))
	}

	for i := 0; i < configFile.RandomNodeTransformationIterations; i++ {
		solutions = append(solutions, transformations.RandomNodeTransformation(data, inputSolution))
	}

	for i := 0; i < configFile.MoveNeighborhoodTransformation; i++ {
		solutions = append(solutions, transformations.MoveNeighborhoodTransformation(data, inputSolution))
	}

	return solutions
}

func GetNeighbourSolution(data *model.CaseDTO, inputSolution *model.Solution, configFile *config.Config) *model.Solution {

	solutions := GetNeighbourSolutions(data, inputSolution, configFile)
	bestSolution := solutions[0]
	for _, solution := range solutions {
		if solution.GetCost(data) < bestSolution.GetCost(data) {
			bestSolution = solution
		}
	}
	return bestSolution
}
