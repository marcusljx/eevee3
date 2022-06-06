package main

import (
	"fmt"
	"github.com/marcusljx/eevee3/lib/eevee3"
	"github.com/marcusljx/eevee3/lib/examples/bestfitline"
	"math/rand"
	"time"
)

var points = []*bestfitline.Point{
	{X: 0, Y: 1},
	{X: 1, Y: 2},
	{X: 2, Y: 3},
	{X: 3, Y: 4},
	{X: 4, Y: 5},
	{X: 5, Y: 6},
	{X: 6, Y: 7},
}

func main() {
	handler := &bestfitline.Handler{
		Rand:   rand.New(rand.NewSource(time.Now().Unix())),
		Points: points,
	}

	data := &eevee3.Controller[bestfitline.TUnderlying]{
		GenerationCycles:                1000,
		PopulationSize:                  10,
		MutationProbability:             1,
		CrossoverSelectionStrategy:      eevee3.SelectRandomPairs[bestfitline.TUnderlying](),
		CrossoverProbability:            0.3,
		NextGenerationSelectionStrategy: eevee3.SelectBestAndWorstSubgroup[bestfitline.TUnderlying](0.8),
		TerminationConditions: []eevee3.PopulationPredicate[bestfitline.TUnderlying]{
			eevee3.TrueWhenAllSolutionsEqual[bestfitline.TUnderlying](),
		},
	}

	result := eevee3.RunSingle[bestfitline.TUnderlying](handler, data)
	fmt.Print(result.Describe())
}
