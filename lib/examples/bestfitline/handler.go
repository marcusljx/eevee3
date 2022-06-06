package bestfitline

import (
	"github.com/marcusljx/eevee3/lib/eevee3"
	"math/rand"
)

type Point struct {
	X float64
	Y float64
}

type Handler struct {
	Rand   *rand.Rand
	Points []*Point
}

func (h *Handler) randomFloat() float64 {
	v := h.Rand.Float64() * float64(h.Rand.Int()%100)
	if h.Rand.Float64() > 0.5 {
		return -v
	}
	return v
}

func (h *Handler) Cross(solution1, solution2 eevee3.Solution[TUnderlying]) (eevee3.Solution[TUnderlying], eevee3.Solution[TUnderlying]) {
	p1, p2 := solution1.Value(), solution2.Value()
	p1[0], p2[0] = p2[0], p1[0]
	return h.NewSolutionFrom(p1), h.NewSolutionFrom(p2)
}

func (h *Handler) Mutate(solution eevee3.Solution[TUnderlying]) eevee3.Solution[TUnderlying] {
	v := solution.Value()
	if h.Rand.Float64() > 0.5 {
		return h.NewSolutionFrom([2]float64{v[0] + h.randomFloat(), v[1]})
	}
	return h.NewSolutionFrom([2]float64{v[0], v[1] + h.randomFloat()})
}

func (h *Handler) NewSolution() eevee3.Solution[TUnderlying] {
	return h.NewSolutionFrom([2]float64{h.randomFloat(), h.randomFloat()})
}

func (h *Handler) NewSolutionFrom(innerValue TUnderlying) eevee3.Solution[TUnderlying] {
	return &Solution{
		polynomial: innerValue,
		handler:    h,
	}
}
