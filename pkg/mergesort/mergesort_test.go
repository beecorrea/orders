package mergesort

import (
	"testing"

	"github.com/beecorrea/orders/pkg/order"
)

func TestMergeSort(t *testing.T) {
	strategy := Mergesort{}
	numbers := order.New([]int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9})
	expected := order.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	actual := numbers.Sort(strategy)
	if len(actual.Members()) != len(expected.Members()) {
		t.Errorf("actual and expected have different amount of elements")
	}

	for i := range actual.Members() {
		if actual.Members()[i] != expected.Members()[i] {
			t.Errorf("actual and expected are in different orders")
		}
	}

	order.AssertPartiallyOrdered(t, numbers)
}
