package order

import (
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
	powerset := set.BuildPowerset(xs)
	return poset{
		members:  xs,
		powerset: powerset,
	}
}

func (ps poset) Members() []int {
	return ps.members
}

func (ps poset) Powerset() set.Powerset {
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
	sets := ps.powerset.FilterBySize(1)
	hasReflexivity := false
	for _, subset := range sets {
		first := subset[0]
		hasReflexivity = Reflexivity(first)
	}

	return hasReflexivity
}

func (ps poset) ensureAntisymmetry() bool {
	sets := ps.powerset.FilterBySize(2)
	hasAntisymmetry := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		hasAntisymmetry = Antisymmetry(first, second)
	}
	return hasAntisymmetry
}

func (ps poset) ensureTransitivity() bool {
	sets := ps.powerset.FilterBySize(3)
	hasTransitivity := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		third := subset[2]
		hasTransitivity = Transitivity(first, second, third)
	}
	return hasTransitivity
}
