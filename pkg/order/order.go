package order

import (
	"fmt"
	"testing"
)

func Reflexivity(a int) bool {
	return a <= a
}

func Antisymmetry(a, b int) bool {
	if (a <= b) && (b <= a) && a != b {
		return false
	}
	return true
}

func Transitivity(a, b, c int) bool {
	if (a <= b) && (b <= c) && !(a <= c) {
		return false
	}

	return true
}

func AssertPartiallyOrdered(t *testing.T) func(strategy Sort, poset, expected Poset) {
	return func(strategy Sort, poset, expected Poset) {
		testName := fmt.Sprintf("Should sort using %s", strategy.Strategy())
		t.Run(testName, func(t *testing.T) {
			actual := poset.Sort(strategy)
			if len(actual.Members()) != len(expected.Members()) {
				t.Errorf("actual and expected have different amount of elements")
			}

			for i := range actual.Members() {
				if actual.Members()[i] != expected.Members()[i] {
					t.Errorf("actual and expected are in different orders")
				}
			}

			t.Run("Assert properties of partial order relations", func(t *testing.T) {
				isPartiallyOrdered := poset.IsPartiallyOrdered()
				if !isPartiallyOrdered {
					t.Errorf("not an poset")
				}
			})
		})
	}
}
