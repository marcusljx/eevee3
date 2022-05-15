package knapsack

import (
	"fmt"
	"strings"
)

const WEIGHT_LIMIT = 1000

// TUnderlying is the underlying type generic of
// a single solution in the experiment corpus
type TUnderlying = []bool

// Solution represents a single possible
// solution instance for the experiment
type Solution struct {
	handler *Handler
	roster  []bool
}

func (s *Solution) getItemAt(i int) Item {
	return s.handler.Items[i]
}

func (s *Solution) String() string {
	buf := sliceMap[bool, byte](s.roster, func(_ int, b bool) byte {
		if b {
			return '1'
		} else {
			return '0'
		}
	})

	return string(buf)
}

func (s *Solution) Value() TUnderlying {
	clone := make([]bool, len(s.roster))
	copy(clone, s.roster)
	return clone
}

// Score for knapsack is the value-per-gram
func (s *Solution) Score() float64 {
	var (
		weightTotal int
		valueTotal  float64
	)
	for i, b := range s.roster {
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
	for i, b := range s.roster {
		if b {
			item := s.getItemAt(i)
			str.WriteString(fmt.Sprintf("- include item %s (weight: %d, value %+0.3f)\n", item.Name, item.WeightInGrams, item.Value))
		}
	}
	return fmt.Sprintf("Score: %+0.3f\nItems:\n%s", s.Score(), str.String())
}
