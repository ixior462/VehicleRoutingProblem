package main

import (
	"fmt"
	"ixior462/VehicleRoutingProblem/algorithms/heuristics"
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/config"
)
const resourcePath = "./resources/X/X-n627-k43.vrp"
const configPath = "./resources/config.json"

func main() {
	inputDTO := common.GetDTO(resourcePath)
	configFile := config.GetConfig(configPath)
	solution := heuristics.TabuSearchHeuristic(inputDTO, configFile)
	//solution := heuristics.LocalSearchHeuristic(inputDTO, configFile)
	//solution := heuristics.SimulatedAnnealingHeuristic(inputDTO, configFile)

	fmt.Println(solution)
}
