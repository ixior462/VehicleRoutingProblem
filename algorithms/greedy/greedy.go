package greedy

import (
	"ixior/VehicleRoutingProblem/common"
	"ixior/VehicleRoutingProblem/model"
)

var depot = 0

func GetInitialSolution(input *model.CaseDTO) *model.Solution {
	solution := model.Solution{}
	unroutedClients := getUnroutedClients(input.GetDimension())
	currentNode := depot
	route := make([]int, 0)

	for len(unroutedClients) > 0 {
		// Search for the best fitting node
		bestNode, bestDistance := unroutedClients[0], input.Cost[currentNode][unroutedClients[0]]
		for i := 1; i < len(unroutedClients); i++ {
			if input.Cost[currentNode][unroutedClients[i]] < bestDistance {
				bestNode, bestDistance = unroutedClients[i], input.Cost[currentNode][unroutedClients[i]]
			}
		}

		// If current route will exceed capacity then create new route
		if common.GetCapacity(route, input)+input.GetDemand(bestNode) > input.Capacity {
			solution.AddRoute(route)
			route = make([]int, 0)
			currentNode = depot
		} else {
			currentNode = bestNode
		}

		// Add the best node to route and then remove it from unroutedClients
		route = append(route, bestNode)
		unroutedClients = removeIndex(unroutedClients, bestNode)
	}

	if len(route) > 0 {
		solution.AddRoute(route)
	}

	return &solution
}

func removeIndex(array []int, element int) []int {
	index := indexOf(element, array)
	return append(array[:index], array[index+1:]...)
}

func getUnroutedClients(dimension int) []int {
	unroutedClients := make([]int, dimension-1)
	for i := 1; i < dimension; i++ {
		unroutedClients[i-1] = i
	}
	return unroutedClients
}

func indexOf(element int, array []int) int {
	for k, v := range array {
		if element == v {
			return k
		}
	}
	return -1
}