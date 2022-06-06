package eevee3

import "fmt"

// Solution is a standard interface for
// a potential solutions in each corpus
type Solution[TUnderlying any] interface {
	fmt.Stringer

	// Score returns the solutions's individual score
	Score() float64

	// Value returns the underlying value backing the solutions
	Value() TUnderlying

	// Equals returns true if the solutions is
	// identical to the given other solutions
	Equals(Solution[TUnderlying]) bool

	// Clone returns a clone that has no shared memory references
	// with the original
	Clone() Solution[TUnderlying]

	// Describe returns textual information describing the solutions.
	// This may be information to be presented at the end of a run, etc.
	Describe() string
}
