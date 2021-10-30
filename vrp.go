package main

import (
	"fmt"
	"ixior462/VehicleRoutingProblem/algorithms/heuristics"
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/config"
)

func main() {
	inputDTO := common.GetDTO("./resources/X/X-n190-k8.vrp")
	configFile := config.GetConfig("./resources/config.json")
	solution := heuristics.TabuSearchHeuristic(inputDTO, configFile)
	//solution := heuristics.LocalSearchHeuristic(inputDTO, configFile)
	//solution := heuristics.SimulatedAnnealingHeuristic(inputDTO, configFile)

	fmt.Println(solution)
}
