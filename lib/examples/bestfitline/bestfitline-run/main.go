package main

import (
	"fmt"
	"github.com/marcusljx/eevee3/lib/eevee3"
	"github.com/marcusljx/eevee3/lib/examples/bestfitline"
	"math/rand"
	"os"
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
	type T = bestfitline.TUnderlying

	handler := &bestfitline.Handler{
		Rand:   rand.New(rand.NewSource(time.Now().Unix())),
		Points: points,
	}

	data := &eevee3.Controller[T]{
		Recorder:                        eevee3.NewIOWriterRecorder[T](os.Stderr, eevee3.RecordEverything),
		GenerationCycles:                1000,
		PopulationSize:                  10,
		MutationProbability:             1,
		CrossoverSelectionStrategy:      eevee3.SelectRandomPairs[T](),
		CrossoverProbability:            0.3,
		NextGenerationSelectionStrategy: eevee3.SelectBestAndWorstSubgroup[T](0.8),
		TerminationConditions: []eevee3.PopulationPredicate[T]{
			eevee3.TrueWhenAllSolutionsEqual[T](),
			eevee3.TrueWhenAtLeastNScores[T](1, func(f float64) bool {
				return f >= 0
			}),
		},
	}

	result := eevee3.RunSingle[T](handler, data)
	fmt.Print(result.Describe())
}
