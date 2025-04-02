package order

import (
	"github.com/beecorrea/orders/pkg/fake"
	"github.com/beecorrea/orders/pkg/set"
)

type Sort interface {
	Strategy() string
	Run(ps Poset) Poset
}

type poset struct {
	members  []int
	powerset set.Powerset
}

type Poset interface {
	Members() []int
	Powerset() set.Powerset
	IsPartiallyOrdered() bool
	Sort(s Sort) Poset
}

func New(xs []int) Poset {
	return poset{
		members:  xs,
		powerset: nil,
	}
}

func Random() Poset {
	xs := fake.RandomInts(10)
	return New(xs)
}

func (ps poset) Members() []int {
	return ps.members
}

func (ps poset) Powerset() set.Powerset {
	if ps.powerset == nil {
		return set.BuildPowerset(ps.members)
	}

	return ps.powerset
}

func (ps poset) IsPartiallyOrdered() bool {
	return ps.ensureReflexivity() &&
		ps.ensureAntisymmetry() &&
		ps.ensureTransitivity()
}

func (ps poset) Sort(s Sort) Poset {
	return s.Run(ps)
}

func (ps poset) ensureReflexivity() bool {
	sets := ps.Powerset().FilterBySize(1)
	hasReflexivity := false
	for _, subset := range sets {
		first := subset[0]
		hasReflexivity = Reflexivity(first)
	}

	return hasReflexivity
}

func (ps poset) ensureAntisymmetry() bool {
	sets := ps.Powerset().FilterBySize(2)
	hasAntisymmetry := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		hasAntisymmetry = Antisymmetry(first, second)
	}
	return hasAntisymmetry
}

func (ps poset) ensureTransitivity() bool {
	sets := ps.Powerset().FilterBySize(3)
	hasTransitivity := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		third := subset[2]
		hasTransitivity = Transitivity(first, second, third)
	}
	return hasTransitivity
}
