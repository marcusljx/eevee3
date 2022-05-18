package eevee3

import "fmt"

// Solution is a standard interface for
// a potential solution in each corpus
type Solution[T any] interface {
	fmt.Stringer

	// Score returns the solution's individual score
	Score() float64

	// Value returns the underlying value backing the solution
	Value() T

	// Equals returns true if the solution is
	// identical to the given other solution
	Equals(Solution[T]) bool

	// Describe returns textual information describing the solution.
	// This may be information to be presented at the end of a run, etc.
	Describe() string
}
