package bestfitline

import (
	"fmt"
	"github.com/marcusljx/eevee3/lib/eevee3"
	"math"
)

// y = mx + c
type linePolynomial [2]float64

func (l linePolynomial) calc(x float64) float64 {
	return (l[0] * x) + l[1]
}

// TUnderlying is the underlying type generic of
// a single solutions in the experiment corpus
type TUnderlying = linePolynomial

// Solution represents a single possible
// solution instance for the experiment
type Solution struct {
	polynomial linePolynomial
	handler    *Handler
}

func (s *Solution) Clone() eevee3.Solution[TUnderlying] {
	return &Solution{
		polynomial: s.polynomial,
		handler:    s.handler,
	}
}

func (s *Solution) String() string {
	return fmt.Sprintf("m = %0.5f\nc = %0.5f", s.polynomial[0], s.polynomial[1])
}

func (s *Solution) Score() (result float64) {
	for _, point := range s.handler.Points {
		result -= math.Abs(point.Y - s.polynomial.calc(point.X))
	}
	return
}

func (s *Solution) Value() TUnderlying {
	return s.polynomial
}

func (s *Solution) Equals(s2 eevee3.Solution[TUnderlying]) bool {
	return s.Value() == s2.Value()
}

func (s *Solution) Describe() string {
	return fmt.Sprintf("Score: %0.3f\n%v", s.Score(), s)
}
