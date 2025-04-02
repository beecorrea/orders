package set

type Powerset [][]int

// Credits to https://prtamil.github.io/posts/powersets-go/
func combinations(xs []int, r int) Powerset {
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
			for _, x := range combinations(perms, r-1) {
				t := append(x, xs[i])
				res = append(res, Powerset{t}...)
			}
		}
		return res
	}
}

// Credits to https://prtamil.github.io/posts/powersets-go/
func BuildPowerset(xs []int) Powerset {
	res := make(Powerset, 0)
	for i := 0; i <= len(xs); i++ {
		x := combinations(xs, i)
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
