package set

import (
	"testing"

	"github.com/beecorrea/orders/pkg/order"
)

type poset struct {
	members  []int
	powerset Powerset
}

type Poset interface {
	Members() []int
	Powerset() Powerset
	IsPartiallyOrdered() bool
}

func New(xs []int) Poset {
	powerset := buildPowerset(xs)
	return poset{
		members:  xs,
		powerset: powerset,
	}
}

func (ps poset) Members() []int {
	return ps.members
}

func (ps poset) Powerset() Powerset {
	return ps.powerset
}

func (ps poset) IsPartiallyOrdered() bool {
	return ps.ensureReflexivity() &&
		ps.ensureAntisymmetry() &&
		ps.ensureTransitivity()
}

func (ps poset) ensureReflexivity() bool {
	sets := ps.powerset.FilterBySize(1)
	hasReflexivity := false
	for _, subset := range sets {
		first := subset[0]
		hasReflexivity = order.Reflexivity(first)
	}

	return hasReflexivity
}

func (ps poset) ensureAntisymmetry() bool {
	sets := ps.powerset.FilterBySize(2)
	hasAntisymmetry := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		hasAntisymmetry = order.Antisymmetry(first, second)
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
		hasTransitivity = order.Transitivity(first, second, third)
	}
	return hasTransitivity
}

func AssertPartiallyOrdered(t *testing.T, poset Poset) bool {
	return t.Run("Assert properties of order relations", func(t *testing.T) {
		isPartiallyOrdered := poset.IsPartiallyOrdered()
		if !isPartiallyOrdered {
			t.Errorf("not an poset")
		}
	})
}
