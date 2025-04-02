package order

import (
	"testing"
)

type OrderTestCases struct {
	poset    Poset
	expected bool
}

func TestIsPartiallyOrdered(t *testing.T) {
	ascendingOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	descendingOrder := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var cases []OrderTestCases = []OrderTestCases{
		{
			poset:    New(ascendingOrder, Leq{}),
			expected: true,
		},
		{
			poset:    New(descendingOrder, Leq{}),
			expected: true,
		},
	}

	for _, c := range cases {
		isOrdered := c.poset.IsPartiallyOrdered()

		if !isOrdered {
			t.Errorf("should be partially ordered")
		}
	}
}

func TestIsntPartiallyOrdered(t *testing.T) {
	ascendingOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	descendingOrder := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var cases []OrderTestCases = []OrderTestCases{
		{
			poset:    New(ascendingOrder, Lt{}),
			expected: true,
		},
		{
			poset:    New(descendingOrder, Lt{}),
			expected: true,
		},
	}

	for _, c := range cases {
		isOrdered := c.poset.IsPartiallyOrdered()

		if !isOrdered {
			t.Errorf("should be partially ordered")
		}
	}
}
