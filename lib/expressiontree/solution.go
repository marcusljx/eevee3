package expressiontree

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
)

type Solution struct {
	root            *Node[float64]
	FitnessFunction func(float64) float64
}

func (s *Solution) String() string {
	return s.Describe()
}

func (s *Solution) Score() float64 {
	return s.FitnessFunction(s.root.Value())
}

func (s *Solution) Value() *Node[float64] {
	return s.root
}

func (s *Solution) Equals(s2 eevee3.Solution[*Node[float64]]) bool {
	return s.Value().Equals(s2.Value())
}

func (s *Solution) Clone() eevee3.Solution[*Node[float64]] {
	//TODO implement me
	panic("implement me")
}

func (s *Solution) Describe() string {
	return s.root.Describe()
}
