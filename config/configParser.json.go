package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	// GetNeighbourSolution
	RandomRouteImprovementTransformationIterations            int     `json:"RandomRouteImprovementTransformationIterations"`
	ShortRouteImprovementTransformationIterations             int     `json:"ShortRouteImprovementTransformationIterations"`
	ReplaceHighestAverageNeighborhoodTransformationIterations int     `json:"ReplaceHighestAverageNeighborhoodTransformationIterations"`
	HighCostNodeImprovementTransformationIterations           int     `json:"HighCostNodeImprovementTransformationIterations"`
	RandomNodeTransformationIterations                        int     `json:"RandomNodeTransformationIterations"`
	MoveNeighborhoodTransformation                            int     `json:"MoveNeighborhoodTransformation"`
	DeterministicNeighborTransformationOccurrenceChance       float64 `json:"DeterministicNeighborTransformationOccurrenceChance"`
	// Heuristics
	LocalSearchHeuristicImprovementCount int     `json:"LocalSearchHeuristicImprovementCount"`
	SimulatedAnnealingAlpha              float64 `json:"SimulatedAnnealingAlpha"`
	SimulatedAnnealingBeta               float64 `json:"SimulatedAnnealingBeta"`
	SimulatedAnnealingM0                 float64 `json:"SimulatedAnnealingM0"`
	SimulatedAnnealingTemperature        float64 `json:"SimulatedAnnealingTemperature"`
	SimulatedAnnealingTemperatureMin     float64 `json:"SimulatedAnnealingTemperatureMin"`
	SimulatedAnnealingTime               float64 `json:"SimulatedAnnealingTime"`
	SimulatedAnnealingTimeMax            float64 `json:"SimulatedAnnealingTimeMax"`
	TabuSearchMaxIterations              int     `json:"TabuSearchMaxIterations"`
	TabuSearchLengthOfTabuList           int     `json:"TabuSearchLengthOfTabuList"`

	// RandomRouteImprovement
	RandomRouteImprovementChanceOfRouteDeletion float64 `json:"RandomRouteImprovementChanceOfRouteDeletion"`
}

func GetConfig(configPath string) *Config {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {

		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Config

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil
	}

	return &users
}
