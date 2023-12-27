package buysell

import (
	"fmt"

	"github.com/marcusljx/eevee3/lib/eevee3"
)

type Action string

const (
	BUY  Action = "buy"
	SELL        = "sell"
	HOLD        = "hold"
)

// TUnderlying is the underlying type generic of
// a single solutions in the experiment corpus
type TUnderlying = []Action

// Solution represents a single possible
// solutions instance for the experiment
type Solution struct {
	eevee3.ConsistentSizeSliceSolution[Action]

	handler *Handler
}

// Score for knapsack is the value-per-gram
func (s *Solution) Score() float64 {
	var (
		shares = 0
		cash   = 0.0
	)

	for i, price := range s.handler.Prices {
		action := s.Slice[i]
		switch action {
		case BUY:
			shares++
			cash -= price

		case SELL:
			shares--
			cash += price

		case HOLD:
			continue
		}
	}

	// after all done, rebalance to get account holdings value at day close
	if shares != 0 {
		cash += float64(shares) * s.handler.Prices[len(s.handler.Prices)-1]
	}

	return cash
}

func (s *Solution) Describe() string {
	return fmt.Sprintf("%0.3f", s.Score())
}
