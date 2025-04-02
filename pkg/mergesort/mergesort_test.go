package mergesort

import (
	"testing"

	"github.com/beecorrea/orders/pkg/order"
)

func TestMergeSort(t *testing.T) {
	strategy := Mergesort{}
	numbers := order.New([]int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9})
	expected := order.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	order.AssertPartiallyOrdered(t)(strategy, numbers, expected)
}
