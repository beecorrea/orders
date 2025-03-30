package set

import "github.com/beecorrea/orders/pkg/order"

type Powerset [][]int

// Credits to https://prtamil.github.io/posts/powersets-go/
func Combinations(xs []int, r int) Powerset {
	if r == 1 {
		temp := make(Powerset, 0)
		for _, rr := range xs {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, Powerset{t}...)
		}
		return temp
	} else {
		res := make(Powerset, 0)
		for i := 0; i < len(xs); i++ {
			perms := make([]int, 0)
			perms = append(perms, xs[:i]...)
			for _, x := range Combinations(perms, r-1) {
				t := append(x, xs[i])
				res = append(res, Powerset{t}...)
			}
		}
		return res
	}
}

// Credits to https://prtamil.github.io/posts/powersets-go/
func GetPowerset(xs []int) Powerset {
	res := make(Powerset, 0)
	for i := 0; i <= len(xs); i++ {
		x := Combinations(xs, i)
		res = append(res, x...)
	}
	return res
}

func (ps Powerset) FilterBySize(size int) Powerset {
	sets := make(Powerset, 0)
	for _, s := range ps {
		if len(s) == size {
			sets = append(sets, s)
		}
	}

	return sets
}

func (ps Powerset) EnsureReflexivity() bool {
	sets := ps.FilterBySize(1)
	hasReflexivity := false
	for _, subset := range sets {
		first := subset[0]
		hasReflexivity = order.Reflexivity(first)
	}

	return hasReflexivity
}

func (ps Powerset) EnsureAntisymmetry() bool {
	sets := ps.FilterBySize(2)
	hasAntisymmetry := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		hasAntisymmetry = order.Antisymmetry(first, second)
	}
	return hasAntisymmetry
}

func (ps Powerset) EnsureTransitivity() bool {
	sets := ps.FilterBySize(3)
	hasTransitivity := false
	for _, subset := range sets {
		first := subset[0]
		second := subset[1]
		third := subset[2]
		hasTransitivity = order.Transitivity(first, second, third)
	}
	return hasTransitivity
}
