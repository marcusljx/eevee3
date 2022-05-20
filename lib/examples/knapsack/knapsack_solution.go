package knapsack

import (
	"fmt"
	"github.com/marcusljx/eevee3/lib/eevee3"
	"strings"
)

const WEIGHT_LIMIT = 1000

// TUnderlying is the underlying type generic of
// a single solutions in the experiment corpus
type TUnderlying = []bool

// Solution represents a single possible
// solutions instance for the experiment
type Solution struct {
	eevee3.ConsistentSizeSliceSolution[bool]

	handler *Handler
}

func (s *Solution) getItemAt(i int) Item {
	return s.handler.Items[i]
}

// Score for knapsack is the value-per-gram
func (s *Solution) Score() float64 {
	var (
		weightTotal int
		valueTotal  float64
	)
	for i, b := range s.Slice {
		if !b {
			continue
		}
		item := s.getItemAt(i)

		weightTotal += item.WeightInGrams
		valueTotal += item.Value
	}

	if weightTotal > WEIGHT_LIMIT {
		return -valueTotal
	}
	return valueTotal
}

func (s *Solution) Describe() string {
	var str strings.Builder
	str.WriteString(fmt.Sprintf("Stringf: %s\n", s.String()))
	for i, b := range s.Slice {
		if b {
			item := s.getItemAt(i)
			str.WriteString(fmt.Sprintf("- include item %s (weight: %d, value %+0.3f)\n", item.Name, item.WeightInGrams, item.Value))
		}
	}
	return fmt.Sprintf("Score: %+0.3f\nItems:\n%s", s.Score(), str.String())
}
