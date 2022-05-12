package eevee3

import "fmt"

type Scoreable interface {
	Score() float64
}

type Solution[T any] interface {
	Scoreable
	fmt.Stringer

	Value() T
	Describe() string
}
