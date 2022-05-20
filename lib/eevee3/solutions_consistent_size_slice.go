package eevee3

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func NewConsistentSizeSliceSolution[E comparable](initialValue []E) *ConsistentSizeSliceSolution[E] {
	return &ConsistentSizeSliceSolution[E]{
		Slice: initialValue,
	}
}

type ConsistentSizeSliceSolution[E comparable] struct {
	Slice []E
}

func (c *ConsistentSizeSliceSolution[E]) String() string {
	return fmt.Sprintf("%v", c.Slice)
}

func (c *ConsistentSizeSliceSolution[E]) Score() float64 {
	panic("Implement Score()")
}

func (c *ConsistentSizeSliceSolution[E]) Value() []E {
	return slices.Clone(c.Slice)
}

func (c *ConsistentSizeSliceSolution[E]) Equals(other Solution[[]E]) bool {
	return slices.Equal(c.Slice, other.Value())
}

func (c *ConsistentSizeSliceSolution[E]) Describe() string {
	return c.String()
}
