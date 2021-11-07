package transformations

import (
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/model"
	"math/rand"
	"sort"
)

// MoveNeighborhoodTransformation remove 5 profitable nodes and randomly relocate them
func MoveNeighborhoodTransformation(data *model.CaseDTO, inputSolution *model.Solution) *model.Solution {
	newSolution := common.CloneSolution(inputSolution)
	unroutedNodes := make([]int, 0)

	for i := 0; i < 5; i++ {
		// Find the most profitable node
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
		if maxI == -1 || maxJ == -1 {
			return inputSolution
		}
		unroutedNodes = append(unroutedNodes, newSolution.Routes[maxI][maxJ])

		// Remove node
		if len(newSolution.Routes[maxI]) == 1 {
			newSolution.Routes = common.RemoveRoute(newSolution.Routes, maxI)
		} else {
			newSolution.Routes[maxI] = common.RemoveNode(newSolution.Routes[maxI], maxJ)
		}
	}

	// sort from most demanding to least
	sort.SliceStable(unroutedNodes, func(i, j int) bool {
		return data.GetDemand(i) > data.GetDemand(j)
	})

	for node := 0; node < 5; node++ {
		if common.CanBePutToRoutes(newSolution, data, unroutedNodes[node]) {
			// Find new place to insert random Node
			for attempts := 0; attempts < 4000; attempts++ {
				if attempts == 3999 {
					return inputSolution
				}
				newRandomRoute := rand.Intn(len(newSolution.Routes))
				if common.GetCapacity(newSolution.Routes[newRandomRoute], data)+data.GetDemand(unroutedNodes[node]) > data.Capacity {
					continue
				}
				newRandomIndex := rand.Intn(len(newSolution.Routes[newRandomRoute]))
				newSolution.Routes[newRandomRoute] = common.AppendAtIndex(newSolution.Routes[newRandomRoute], newRandomIndex, unroutedNodes[node])
				break
			}
		} else {
			newSolution.AddRoute([]int{unroutedNodes[node]})
		}
	}

	return newSolution
}
