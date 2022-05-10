package main

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
	"github.com/marcusljx/eevee3/lib/examples/knapsack"
	"math/rand"
	"time"
)

var items = []knapsack.Item{
	{WeightInGrams: 100, Value: 5},
}

func main() {
	knapsackHandler := &knapsack.Handler{
		Rand:  rand.New(rand.NewSource(time.Now().Unix())),
		Items: items,
	}

	data := &eevee3.ExperimentData[knapsack.SolutionType]{
		Generations:                     100,
		PopulationSize:                  10,
		MutationProbability:             0.05,
		CrossoverSelectionStrategy:      eevee3.SelectRandomPairs[knapsack.SolutionType],
		CrossoverProbability:            0,
		NextGenerationSelectionStrategy: eevee3.SelectBestSubgroup[knapsack.SolutionType],
	}

	eevee3.Run[knapsack.SolutionType](knapsackHandler, data)
}
