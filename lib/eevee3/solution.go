package eevee3

import "fmt"

// Solution is a standard interface for
// a potential solutions in each corpus
type Solution[T any] interface {
	fmt.Stringer

	// Score returns the solutions's individual score
	Score() float64

	// Value returns the underlying value backing the solutions
	Value() T

	// Equals returns true if the solutions is
	// identical to the given other solutions
	Equals(Solution[T]) bool

	// Describe returns textual information describing the solutions.
	// This may be information to be presented at the end of a run, etc.
	Describe() string
}
