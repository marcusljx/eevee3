package eevee3

import (
	"fmt"
	"strings"
)

// Result is the output of a single experiment run
type Result[T any] struct {
	Population             []Solution[T] `json:"population"`
	TerminatedAtGeneration int           `json:"terminatedAtGeneration"`
	TerminationReason      string        `json:"terminationReason"`
}

func (r *Result[T]) Describe() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Terminated at Generation: %d\n", r.TerminatedAtGeneration))
	builder.WriteString(fmt.Sprintf("Reason for Termination:   %s\n", r.TerminationReason))
	builder.WriteString("Population Description(s):\n")
	for _, sol := range r.Population {
		builder.WriteString(sol.Describe())
	}
	return builder.String()
}
