package main

import (
	"fmt"
	"github.com/marcusljx/eevee3/lib/eevee3"
	"github.com/marcusljx/eevee3/lib/examples/knapsack"
	"math/rand"
	"time"
)

var items = []knapsack.Item{
	{Name: "A", WeightInGrams: 500, Value: 40},
	{Name: "B", WeightInGrams: 20, Value: 50},
	{Name: "C", WeightInGrams: 100, Value: 50},
	{Name: "D", WeightInGrams: 3, Value: 17},
	{Name: "E", WeightInGrams: 100, Value: 140},
	{Name: "F", WeightInGrams: 800, Value: 290},
	{Name: "G", WeightInGrams: 10, Value: 80},
}

func main() {
	knapsackHandler := &knapsack.Handler{
		Rand:  rand.New(rand.NewSource(time.Now().Unix())),
		Items: items,
	}

	data := &eevee3.ExperimentData[knapsack.TUnderlying]{
		Generations:                     100,
		PopulationSize:                  10,
		MutationProbability:             0.05,
		CrossoverSelectionStrategy:      eevee3.SelectRandomPairs[knapsack.TUnderlying],
		CrossoverProbability:            0,
		NextGenerationSelectionStrategy: eevee3.SelectBestSubgroup[knapsack.TUnderlying],
	}

	result := eevee3.Run[knapsack.TUnderlying](knapsackHandler, data)

	for _, sol := range result {
		fmt.Printf("(%0.4f)%s\n", sol.Score(), sol)
	}
}
