package transformations

import (
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/model"
	"math/rand"
)

// RandomNodeTransformation Get random node from solutions, and relocate it randomly
func RandomNodeTransformation(data *model.CaseDTO, inputSolution *model.Solution) *model.Solution {
	newSolution := common.CloneSolution(inputSolution)

	// Remove random node
	randomRoute := rand.Intn(len(inputSolution.Routes))
	randomIndex := rand.Intn(len(inputSolution.Routes[randomRoute]))
	removedNode := newSolution.Routes[randomRoute][randomIndex]
	newSolution.Routes[randomRoute] = common.RemoveNode(newSolution.Routes[randomRoute], randomIndex)

	// If route is now empty, remove it
	if len(newSolution.Routes[randomRoute]) == 0 {
		newSolution.Routes = common.RemoveRoute(newSolution.Routes, randomRoute)
	}

	// Find new place to insert random Node
	for attempts := 0; attempts < 50; attempts++ {
		newRandomRoute := rand.Intn(len(newSolution.Routes))
		if common.GetCapacity(newSolution.Routes[newRandomRoute], data)+data.GetDemand(removedNode) > data.Capacity {
			continue
		}
		newRandomIndex := rand.Intn(len(newSolution.Routes[newRandomRoute]))
		newSolution.Routes[newRandomRoute] = common.AppendAtIndex(newSolution.Routes[newRandomRoute], newRandomIndex, removedNode)
		return newSolution
	}

	// If random insertion failed, just return original solution
	return inputSolution
}
