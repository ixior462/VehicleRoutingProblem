package transformations

import (
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/model"
	"math"
)

// HighCostNodeImprovementTransformation Remove 1 most unprofitable node and then find optimal place for it
func HighCostNodeImprovementTransformation(data *model.CaseDTO, inputSolution *model.Solution) *model.Solution {
	newSolution := common.CloneSolution(inputSolution)

	// Find the most unprofitable node
	highestCost := 0.0
	maxI, maxJ := -1, -1
	for i := 0; i < len(newSolution.Routes); i++ {
		for j := 0; j < len(newSolution.Routes[i]); j++ {
			node1, node2, node3 := 0, newSolution.Routes[i][j], 0
			if j != 0 {
				node1 = newSolution.Routes[i][j-1]
			}
			if j != len(newSolution.Routes[i])-1 {
				node3 = newSolution.Routes[i][j+1]
			}
			cost := data.Cost[node1][node2] + data.Cost[node2][node3]
			if cost > highestCost {
				highestCost, maxI, maxJ = cost, i, j
			}
		}
	}

	removedNode := newSolution.Routes[maxI][maxJ]

	// Remove node
	if len(newSolution.Routes[maxI]) == 1 {
		newSolution.Routes = common.RemoveRoute(newSolution.Routes, maxI)
	} else {
		newSolution.Routes[maxI] = common.RemoveNode(newSolution.Routes[maxI], maxJ)
	}

	// Find new place for node
	lowestCost := math.MaxFloat64
	minI, minJ := -1, -1
	for i := 0; i < len(newSolution.Routes); i++ {
		if common.GetCapacity(newSolution.Routes[i], data)+data.GetDemand(removedNode) <= data.Capacity {
			for j := 0; j <= len(newSolution.Routes[i]); j++ {
				node1, node2 := 0, 0
				if j != 0 {
					node1 = newSolution.Routes[i][j-1]
				}
				if j != len(newSolution.Routes[i]) {
					node2 = newSolution.Routes[i][j]
				}
				cost := data.Cost[node1][removedNode] + data.Cost[removedNode][node2]
				if cost < lowestCost {
					lowestCost, minI, minJ = cost, i, j
				}
			}
		}
	}

	if minI == -1 || minJ == -1 {
		return inputSolution
	}
	// Insert removed node
	newSolution.Routes[minI] = common.AppendAtIndex(newSolution.Routes[minI], minJ, removedNode)
	return newSolution
}
