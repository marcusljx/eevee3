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

	data := &eevee3.Controller[knapsack.TUnderlying]{
		GenerationCycles:                100,
		PopulationSize:                  10,
		MutationProbability:             0.05,
		CrossoverSelectionStrategy:      eevee3.SelectRandomPairs[knapsack.TUnderlying](),
		CrossoverProbability:            1,
		NextGenerationSelectionStrategy: eevee3.SelectBestAndWorstSubgroup[knapsack.TUnderlying](0.8),
		TerminationConditions: []eevee3.PopulationPredicate[knapsack.TUnderlying]{
			eevee3.TrueWhenAllSolutionsEqual[knapsack.TUnderlying](),
		},
	}

	result := eevee3.RunSingle[knapsack.TUnderlying](knapsackHandler, data)
	fmt.Print(result.Describe())
}
