package mergesort

import (
	"testing"

	"github.com/beecorrea/orders/pkg/order"
)

func TestMergeSort(t *testing.T) {
	strategy := Mergesort{}
	numbers := order.Random()
	// t.Log(numbers)
	order.AssertPartiallyOrdered(t)(strategy, numbers)
}
