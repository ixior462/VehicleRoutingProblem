package transformations

import (
	"ixior462/VehicleRoutingProblem/algorithms/greedy"
	"ixior462/VehicleRoutingProblem/common"
	"ixior462/VehicleRoutingProblem/config"
	"ixior462/VehicleRoutingProblem/model"
	"testing"
)

func TestTransformations(t *testing.T) {
	list := [...]string{
		"./../../resources/A/A-n32-k5.vrp",
		"./../../resources/A/A-n33-k5.vrp",
		"./../../resources/A/A-n33-k6.vrp",
		"./../../resources/A/A-n34-k5.vrp",
		"./../../resources/A/A-n36-k5.vrp",
		"./../../resources/A/A-n37-k5.vrp",
		"./../../resources/A/A-n37-k6.vrp",
		"./../../resources/A/A-n38-k5.vrp",
		"./../../resources/A/A-n39-k5.vrp",
		"./../../resources/A/A-n39-k6.vrp",
		"./../../resources/A/A-n44-k6.vrp",
		"./../../resources/A/A-n45-k6.vrp",
		"./../../resources/A/A-n45-k7.vrp",
		"./../../resources/A/A-n46-k7.vrp",
		"./../../resources/A/A-n48-k7.vrp",
		"./../../resources/A/A-n53-k7.vrp",
		"./../../resources/A/A-n54-k7.vrp",
		"./../../resources/A/A-n55-k9.vrp",
		"./../../resources/A/A-n60-k9.vrp",
		"./../../resources/A/A-n61-k9.vrp",
		"./../../resources/A/A-n62-k8.vrp",
		"./../../resources/A/A-n63-k10.vrp",
		"./../../resources/A/A-n63-k9.vrp",
		"./../../resources/A/A-n64-k9.vrp",
		"./../../resources/A/A-n65-k9.vrp",
		"./../../resources/A/A-n69-k9.vrp",
		"./../../resources/A/A-n80-k10.vrp",
		"./../../resources/B/B-n31-k5.vrp",
		"./../../resources/B/B-n34-k5.vrp",
		"./../../resources/B/B-n35-k5.vrp",
		"./../../resources/B/B-n38-k6.vrp",
		"./../../resources/B/B-n39-k5.vrp",
		"./../../resources/B/B-n41-k6.vrp",
		"./../../resources/B/B-n43-k6.vrp",
		"./../../resources/B/B-n44-k7.vrp",
		"./../../resources/B/B-n45-k5.vrp",
		"./../../resources/B/B-n45-k6.vrp",
		"./../../resources/B/B-n50-k7.vrp",
		"./../../resources/B/B-n50-k8.vrp",
		"./../../resources/B/B-n51-k7.vrp",
		"./../../resources/B/B-n52-k7.vrp",
		"./../../resources/B/B-n56-k7.vrp",
		"./../../resources/B/B-n57-k7.vrp",
		"./../../resources/B/B-n57-k9.vrp",
		"./../../resources/B/B-n63-k10.vrp",
		"./../../resources/B/B-n64-k9.vrp",
		"./../../resources/B/B-n66-k9.vrp",
		"./../../resources/B/B-n67-k10.vrp",
		"./../../resources/B/B-n68-k9.vrp",
		"./../../resources/B/B-n78-k10.vrp",
		"./../../resources/F/F-n135-k7.vrp",
		"./../../resources/F/F-n45-k4.vrp",
		"./../../resources/F/F-n72-k4.vrp",
		"./../../resources/M/M-n101-k10.vrp",
		"./../../resources/M/M-n121-k7.vrp",
		"./../../resources/M/M-n151-k12.vrp",
		"./../../resources/M/M-n200-k16.vrp",
		"./../../resources/M/M-n200-k17.vrp",
		"./../../resources/P/P-n101-k4.vrp",
		"./../../resources/P/P-n16-k8.vrp",
		"./../../resources/P/P-n19-k2.vrp",
		"./../../resources/P/P-n20-k2.vrp",
		"./../../resources/P/P-n21-k2.vrp",
		"./../../resources/P/P-n22-k2.vrp",
		"./../../resources/P/P-n22-k8.vrp",
		"./../../resources/P/P-n23-k8.vrp",
		"./../../resources/P/P-n40-k5.vrp",
		"./../../resources/P/P-n45-k5.vrp",
		"./../../resources/P/P-n50-k10.vrp",
		"./../../resources/P/P-n50-k7.vrp",
		"./../../resources/P/P-n50-k8.vrp",
		"./../../resources/P/P-n51-k10.vrp",
		"./../../resources/P/P-n55-k10.vrp",
		"./../../resources/P/P-n55-k15.vrp",
		"./../../resources/P/P-n55-k7.vrp",
		"./../../resources/P/P-n55-k8.vrp",
		"./../../resources/P/P-n60-k10.vrp",
		"./../../resources/P/P-n60-k15.vrp",
		"./../../resources/P/P-n65-k10.vrp",
		"./../../resources/P/P-n70-k10.vrp",
		"./../../resources/P/P-n76-k4.vrp",
		"./../../resources/P/P-n76-k5.vrp",
		"./../../resources/X/X-n1001-k43.vrp",
		"./../../resources/X/X-n101-k25.vrp",
		"./../../resources/X/X-n106-k14.vrp",
		"./../../resources/X/X-n110-k13.vrp",
		"./../../resources/X/X-n115-k10.vrp",
		"./../../resources/X/X-n120-k6.vrp",
		"./../../resources/X/X-n125-k30.vrp",
		"./../../resources/X/X-n129-k18.vrp",
		"./../../resources/X/X-n134-k13.vrp",
		"./../../resources/X/X-n139-k10.vrp",
		"./../../resources/X/X-n143-k7.vrp",
		"./../../resources/X/X-n148-k46.vrp",
		"./../../resources/X/X-n153-k22.vrp",
		"./../../resources/X/X-n157-k13.vrp",
		"./../../resources/X/X-n162-k11.vrp",
		"./../../resources/X/X-n167-k10.vrp",
		"./../../resources/X/X-n172-k51.vrp",
		"./../../resources/X/X-n176-k26.vrp",
		"./../../resources/X/X-n181-k23.vrp",
		"./../../resources/X/X-n186-k15.vrp",
		"./../../resources/X/X-n190-k8.vrp",
		"./../../resources/X/X-n195-k51.vrp",
		"./../../resources/X/X-n200-k36.vrp",
		"./../../resources/X/X-n204-k19.vrp",
		"./../../resources/X/X-n209-k16.vrp",
		"./../../resources/X/X-n214-k11.vrp",
		"./../../resources/X/X-n219-k73.vrp",
		"./../../resources/X/X-n223-k34.vrp",
		"./../../resources/X/X-n228-k23.vrp",
		"./../../resources/X/X-n233-k16.vrp",
		"./../../resources/X/X-n237-k14.vrp",
		"./../../resources/X/X-n242-k48.vrp",
		"./../../resources/X/X-n247-k50.vrp",
		"./../../resources/X/X-n251-k28.vrp",
		"./../../resources/X/X-n256-k16.vrp",
		"./../../resources/X/X-n261-k13.vrp",
		"./../../resources/X/X-n266-k58.vrp",
		"./../../resources/X/X-n270-k35.vrp",
		"./../../resources/X/X-n275-k28.vrp",
		"./../../resources/X/X-n280-k17.vrp",
		"./../../resources/X/X-n284-k15.vrp",
		"./../../resources/X/X-n289-k60.vrp",
		"./../../resources/X/X-n294-k50.vrp",
		"./../../resources/X/X-n298-k31.vrp",
		"./../../resources/X/X-n303-k21.vrp",
		"./../../resources/X/X-n308-k13.vrp",
		"./../../resources/X/X-n313-k71.vrp",
		"./../../resources/X/X-n317-k53.vrp",
		"./../../resources/X/X-n322-k28.vrp",
		"./../../resources/X/X-n327-k20.vrp",
		"./../../resources/X/X-n331-k15.vrp",
		"./../../resources/X/X-n336-k84.vrp",
		"./../../resources/X/X-n344-k43.vrp",
		"./../../resources/X/X-n351-k40.vrp",
		"./../../resources/X/X-n359-k29.vrp",
		"./../../resources/X/X-n367-k17.vrp",
		"./../../resources/X/X-n376-k94.vrp",
		"./../../resources/X/X-n384-k52.vrp",
		"./../../resources/X/X-n393-k38.vrp",
		"./../../resources/X/X-n401-k29.vrp",
		"./../../resources/X/X-n411-k19.vrp",
		"./../../resources/X/X-n420-k130.vrp",
		"./../../resources/X/X-n429-k61.vrp",
		"./../../resources/X/X-n439-k37.vrp",
		"./../../resources/X/X-n449-k29.vrp",
		"./../../resources/X/X-n459-k26.vrp",
		"./../../resources/X/X-n469-k138.vrp",
		"./../../resources/X/X-n480-k70.vrp",
		"./../../resources/X/X-n491-k59.vrp",
		"./../../resources/X/X-n502-k39.vrp",
		"./../../resources/X/X-n513-k21.vrp",
		"./../../resources/X/X-n524-k153.vrp",
		"./../../resources/X/X-n536-k96.vrp",
		"./../../resources/X/X-n548-k50.vrp",
		"./../../resources/X/X-n561-k42.vrp",
		"./../../resources/X/X-n573-k30.vrp",
		"./../../resources/X/X-n586-k159.vrp",
		"./../../resources/X/X-n599-k92.vrp",
		"./../../resources/X/X-n613-k62.vrp",
		"./../../resources/X/X-n627-k43.vrp",
		"./../../resources/X/X-n641-k35.vrp",
		"./../../resources/X/X-n655-k131.vrp",
		"./../../resources/X/X-n670-k130.vrp",
		"./../../resources/X/X-n685-k75.vrp",
		"./../../resources/X/X-n701-k44.vrp",
		"./../../resources/X/X-n716-k35.vrp",
		"./../../resources/X/X-n733-k159.vrp",
		"./../../resources/X/X-n749-k98.vrp",
		"./../../resources/X/X-n766-k71.vrp",
		"./../../resources/X/X-n783-k48.vrp",
		"./../../resources/X/X-n801-k40.vrp",
		"./../../resources/X/X-n819-k171.vrp",
		"./../../resources/X/X-n837-k142.vrp",
		"./../../resources/X/X-n856-k95.vrp",
		"./../../resources/X/X-n876-k59.vrp",
		"./../../resources/X/X-n895-k37.vrp",
		"./../../resources/X/X-n916-k207.vrp",
		"./../../resources/X/X-n936-k151.vrp",
		"./../../resources/X/X-n957-k87.vrp",
		//"./../../resources/XXL/Antwerp1.vrp",
		//"./../../resources/XXL/Antwerp2.vrp",
		//"./../../resources/XXL/Brussels1.vrp",
		//"./../../resources/XXL/Brussels2.vrp",
		//"./../../resources/XXL/Flanders1.vrp",
		//"./../../resources/XXL/Flanders2.vrp",
		//"./../../resources/XXL/Ghent1.vrp",
		//"./../../resources/XXL/Ghent2.vrp",
		//"./../../resources/XXL/Leuven1.vrp",
		//"./../../resources/XXL/Leuven2.vrp",
		"./../../resources/X/X-n979-k58.vrp"}
	configFile := config.GetConfig("./../../resources/config.json")
	for _, example := range list {
		inputDTO := common.GetDTO(example)
		initialSolution := greedy.GetInitialSolution(inputDTO)
		if checkSolution(inputDTO, initialSolution) {
			t.Error("initialSolution failed for " + example)
		}
		solution1 := RandomNodeTransformation(inputDTO, initialSolution)
		if checkSolution(inputDTO, solution1) {
			t.Error("RandomNodeTransformation failed for " + example)
		}
		solution2 := HighCostNodeImprovementTransformation(inputDTO, initialSolution)
		if checkSolution(inputDTO, solution2) {
			t.Error("HighCostNodeImprovementTransformation failed for " + example)
		}
		solution3 := ShortRouteImprovementTransformation(inputDTO, initialSolution)
		if checkSolution(inputDTO, solution3) {
			t.Error("ShortRouteImprovementTransformation failed for " + example)
		}
		solution4 := RandomRouteImprovementTransformation(inputDTO, initialSolution, configFile)
		if checkSolution(inputDTO, solution4) {
			t.Error("RandomRouteImprovementTransformation failed for " + example)
		}
		solution5 := ReplaceHighestAverageNeighborhoodTransformation(inputDTO, initialSolution)
		if checkSolution(inputDTO, solution5) {
			t.Error("ReplaceHighestAverageNeighborhoodTransformation failed for " + example)
		}
		solution6 := MoveNeighborhoodTransformation(inputDTO, initialSolution)
		if checkSolution(inputDTO, solution6) {
			t.Error("ReplaceHighestAverageNeighborhoodTransformation failed for " + example)
		}
	}

}

func checkSolution(data *model.CaseDTO, solution *model.Solution) bool {
	sum := 0
	set := make(map[int]bool)
	for i := 0; i < len(solution.Routes); i++ {
		sum += len(solution.Routes[i])
		if common.GetCapacity(solution.Routes[i], data) > data.Capacity {
			return true
		}
		for j := 0; j < len(solution.Routes[i]); j++ {
			set[solution.Routes[i][j]] = true
		}
	}

	if sum != data.Dimension-1 && len(set) != sum {
		return true
	}
	return false
}
