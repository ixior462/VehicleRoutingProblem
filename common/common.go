package common

import (
	"bufio"
	"ixior462/VehicleRoutingProblem/model"
	"log"
	"math/rand"
	"os"
	"time"
)

func GetCapacity(route []int, data *model.CaseDTO) int {
	capacity := 0
	for _, node := range route {
		capacity += data.GetDemand(node)
	}
	return capacity
}

func GetDTO(path string) *model.CaseDTO {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	data := model.GetCaseDTO(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func CloneSolution(solution *model.Solution) *model.Solution {
	copiedRoutes := make([][]int, len(solution.Routes))

	for i := 0; i < len(copiedRoutes); i++ {
		copiedRoutes[i] = make([]int, len(solution.Routes[i]))
		copy(copiedRoutes[i], solution.Routes[i])
	}

	return &model.Solution{Routes: copiedRoutes}
}

func CanBePutToRoutes(solution *model.Solution, data *model.CaseDTO, node int) bool {
	for i := 0; i < len(solution.Routes); i++ {
		if GetCapacity(solution.Routes[i], data)+data.GetDemand(node) <= data.Capacity {
			return true
		}
	}
	return false
}

func Shuffle(nodes []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
}

func AppendAtIndex(slice []int, index int, value int) []int {
	if len(slice) == index {
		slice = append(slice, value)
		return slice
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func RemoveNode(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveRoute(slice [][]int, index int) [][]int {
	return append(slice[:index], slice[index+1:]...)
}
