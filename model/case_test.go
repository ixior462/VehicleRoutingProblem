package model

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestParsingSourceFile(t *testing.T) {
	f, _ := os.Open("./../resources/A/A-n32-k5.vrp")
	scanner := bufio.NewScanner(f)

	dto := GetCaseDTO(scanner)

	if dto == nil {
		t.Error("Expected Case DTO to not be null")
	}
	if dto.name != "A-n32-k5" {
		t.Error("Unexpected name: " + dto.name)
	}
	if dto.comment != "(Augerat et al, No of trucks: 5, Optimal value: 784)" {
		t.Error("Unexpected comment: " + dto.comment)
	}
	if dto.vrpType != "CVRP" {
		t.Error("Unexpected vrpType: " + dto.vrpType)
	}
	if dto.Dimension != 32 {
		t.Error("Unexpected dimension: " + strconv.Itoa(dto.Dimension))
	}
	if dto.vrpType != "CVRP" {
		t.Error("Unexpected vrpType: " + dto.vrpType)
	}
	if dto.edgeWeightType != "EUC_2D" {
		t.Error("Unexpected vrpType: " + dto.vrpType)
	}
	if dto.vrpType != "CVRP" {
		t.Error("Unexpected vrpType: " + dto.vrpType)
	}
	if dto.Capacity != 100 {
		t.Error("Unexpected capacity: " + strconv.Itoa(dto.Capacity))
	}
	if dto.Depot != 1 {
		t.Error("Unexpected depot: " + strconv.Itoa(dto.Depot))
	}

	expectedDemands := [32]int{0, 19, 21, 6, 19, 7, 12, 16, 6, 16, 8, 14, 21, 16,
		3, 22, 18, 19, 1, 24, 8, 12, 4, 8, 24, 24, 2, 20, 15, 2, 14, 9}

	for i := 0; i < 32; i++ {
		if dto.Demands[i] != expectedDemands[i] {
			t.Error("Unexpected demand found: " + strconv.Itoa(i) +
				" but  it should be " + strconv.Itoa(expectedDemands[i]) + "at index: " + strconv.Itoa(i))
		}
	}

	expectedNodes := [32][2]float64{
		{82, 76},
		{96, 44},
		{50, 5},
		{49, 8},
		{13, 7},
		{29, 89},
		{58, 30},
		{84, 39},
		{14, 24},
		{2, 39},
		{3, 82},
		{5, 10},
		{98, 52},
		{84, 25},
		{61, 59},
		{1, 65},
		{88, 51},
		{91, 2},
		{19, 32},
		{93, 3},
		{50, 93},
		{98, 14},
		{5, 42},
		{42, 9},
		{61, 62},
		{9, 97},
		{80, 55},
		{57, 69},
		{23, 15},
		{20, 70},
		{85, 60},
		{98, 5}}

	for i := 0; i < 32; i++ {
		if dto.nodes[i][0] != expectedNodes[i][0] || dto.nodes[i][1] != expectedNodes[i][1] {
			t.Error("Unexpected node found at index: " + strconv.Itoa(i))
		}
	}
}
