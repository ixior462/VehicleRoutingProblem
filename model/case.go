package model

import (
	"bufio"
	"log"
	"math"
	"strconv"
	"strings"
)

type CaseDTO struct {
	name           string
	comment        string
	vrpType        string
	Dimension      int
	edgeWeightType string
	Capacity       int
	Depot          int
	Demands        []int
	nodes          [][]float64
	Cost           [][]float64
}

func GetCaseDTO(scanner *bufio.Scanner) *CaseDTO {
	dto := CaseDTO{}
	flag := ""
	row := 0

	for scanner.Scan() {
		line := scanner.Text()

		// ignore empty lines
		if len(line) == 0 {
			continue
		}

		// get rid off whitespaces
		line = strings.TrimSpace(line)

		// split the line
		words := strings.Fields(line)
		if flag != "" {
			if _, err := strconv.Atoi(words[0]); err != nil {
				flag = ""
				row = 0
			} else {
				switch flag {
				case "NODE_COORD_SECTION":
					x, xErr := strconv.ParseFloat(words[1], 64)
					y, yErr := strconv.ParseFloat(words[2], 64)
					if xErr != nil && yErr != nil {
						log.Fatal("Error while loading demands")
					}
					if dto.nodes == nil {
						dto.nodes = make([][]float64, dto.Dimension)
					}
					dto.nodes[row] = make([]float64, 2)
					dto.nodes[row][0] = x
					dto.nodes[row][1] = y
				case "DEMAND_SECTION":
					_, nodeErr := strconv.Atoi(words[0])
					demand, demandErr := strconv.Atoi(words[1])
					if nodeErr != nil && demandErr != nil {
						log.Fatal("Error while loading demands")
					}
					if dto.Demands == nil {
						dto.Demands = make([]int, dto.Dimension)
					}
					dto.Demands[row] = demand
				case "DEPOT_SECTION":
					depot, _ := strconv.Atoi(words[0])
					if depot != -1 {
						dto.Depot = depot
					}
				}
				row++
			}
		}

		switch words[0] {
		case "NAME":
			dto.name = words[2]
		case "COMMENT":
			dto.comment = strings.Replace(line, "COMMENT : ", "", 1)
		case "TYPE":
			dto.vrpType = words[2]
		case "DIMENSION":
			word, err := strconv.Atoi(words[2])
			if err != nil {
				log.Fatal(err)
			}
			dto.Dimension = word
		case "EDGE_WEIGHT_TYPE":
			dto.edgeWeightType = words[2]
		case "CAPACITY":
			word, err := strconv.Atoi(words[2])
			if err != nil {
				log.Fatal(err)
			}
			dto.Capacity = word
		case "NODE_COORD_SECTION":
			flag = "NODE_COORD_SECTION"
		case "DEMAND_SECTION":
			flag = "DEMAND_SECTION"
		case "DEPOT_SECTION":
			flag = "DEPOT_SECTION"
		case "EOF":
			continue
		case "-1":
			continue
		}
	}

	// calculate costs
	dto.Cost = make([][]float64, dto.Dimension)
	for i := range dto.Cost {
		dto.Cost[i] = make([]float64, dto.Dimension)
	}

	for i := 0; i < dto.Dimension-1; i++ {
		x := dto.nodes[i][0]
		y := dto.nodes[i][1]
		for j := i + 1; j < dto.Dimension; j++ {
			value := distance(x, dto.nodes[j][0], y, dto.nodes[j][1])
			dto.Cost[i][j] = value
			dto.Cost[j][i] = value
		}
	}
	return &dto
}

func distance(x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	first := math.Pow(float64(x1-x2), 2)
	second := math.Pow(float64(y1-y2), 2)
	return math.Sqrt(first + second)
}

func (caseDTO *CaseDTO) GetCosts() [][]float64 {
	return caseDTO.Cost
}

func (caseDTO *CaseDTO) GetDemand(i int) int {
	return caseDTO.Demands[i]
}

func (caseDTO *CaseDTO) GetName() string {
	return caseDTO.name
}

func (caseDTO *CaseDTO) GetVrpType() string {
	return caseDTO.vrpType
}

func (caseDTO *CaseDTO) GetDimension() int {
	return caseDTO.Dimension
}

func (caseDTO *CaseDTO) GetEdgeWeightType() string {
	return caseDTO.edgeWeightType
}

func (caseDTO *CaseDTO) GetCapacity() int {
	return caseDTO.Capacity
}

func (caseDTO *CaseDTO) GetComment() string {
	return caseDTO.comment
}

func (caseDTO *CaseDTO) GetDepot() int {
	return caseDTO.Depot
}
