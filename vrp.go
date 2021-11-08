package main

import (
	"fmt"
	"ixior462/VehicleRoutingProblem/algorithms/heuristics"
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/config"
)

const resourcePath = "./resources/M/M-n200-k17.vrp"
const configPath = "./resources/config.json"

func main() {
	data := common.GetDTO(resourcePath)
	configFile := config.GetConfig(configPath)
	tabuSearchSolution := heuristics.TabuSearchHeuristic(data, configFile)
	localSearchSolution := heuristics.LocalSearchHeuristic(data, configFile)
	simulatedAnnealingSolution := heuristics.SimulatedAnnealingHeuristic(data, configFile)

	fmt.Println("LocalSearchHeuristic: ", localSearchSolution.GetCost(data))
	fmt.Println("SimulatedAnnealingHeuristic: ", simulatedAnnealingSolution.GetCost(data))
	fmt.Println("TabuSearchHeuristic: ", tabuSearchSolution.GetCost(data))
}
