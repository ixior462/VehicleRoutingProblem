package heuristics

import (
	"ixior462/VehicleRoutingProblem/algorithms"
	"ixior462/VehicleRoutingProblem/algorithms/greedy"
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/config"
	"ixior462/VehicleRoutingProblem/model"
	"math"
	"math/rand"
)

func SimulatedAnnealingHeuristic(data *model.CaseDTO, configFile *config.Config) *model.Solution {

	currentSolution := greedy.GetInitialSolution(data)
	bestSolution := common.CloneSolution(currentSolution)

	//Temperature reduction multiplier
	alpha := configFile.SimulatedAnnealingAlpha

	// Iteration multiplier
	beta := configFile.SimulatedAnnealingBeta

	//Time until next parameter update
	M0 := configFile.SimulatedAnnealingM0

	temperature := configFile.SimulatedAnnealingTemperature
	temperatureMin := configFile.SimulatedAnnealingTemperatureMin
	time := configFile.SimulatedAnnealingTime
	timeMax := configFile.SimulatedAnnealingTimeMax

	for time < timeMax && temperature > temperatureMin {
		M := M0
		for M >= 0 {
			newSolution := algorithms.GetNeighbourSolution(data, currentSolution, configFile)
			delta := newSolution.GetCost(data) - currentSolution.GetCost(data)
			if delta < 0.0 {
				currentSolution = newSolution
				if currentSolution.GetCost(data) < bestSolution.GetCost(data) {
					bestSolution = currentSolution
				}
			} else if rand.Float64() < math.Exp(-delta/temperature) {
				currentSolution = newSolution
			}
			M--
		}
		time = time + M0
		temperature = alpha * temperature
		M0 = beta * M0
	}

	return bestSolution
}
