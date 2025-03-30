package mergesort

import (
	"testing"

	"github.com/beecorrea/orders/pkg/set"
)

func TestMergeSort(t *testing.T) {
	numbers := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	actual := Mergesort(numbers)
	if len(actual) != len(expected) {
		t.Errorf("actual and expected have different amount of elements")
	}

	i := 0
	for i < len(actual) {
		if actual[i] != expected[i] {
			t.Errorf("actual and expected are in different orders")
		}
		i++
	}

	t.Run("Ensure properties of order relations", func(t *testing.T) {
		powerset := set.GetPowerset(numbers)
		hasReflexivity := powerset.EnsureReflexivity()
		hasAntisymmetry := powerset.EnsureAntisymmetry()
		hasTransitivity := powerset.EnsureTransitivity()

		if !(hasReflexivity && hasAntisymmetry && hasTransitivity) {
			t.Errorf("not an order")
		}
	})
}
