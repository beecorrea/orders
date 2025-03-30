package mergesort

func sortAndMerge(xs []int, ys []int) []int {
	i := 0
	j := 0
	out := make([]int, 0)

	for i < len(xs) && j < len(ys) {
		if xs[i] < ys[j] {
			out = append(out, xs[i])
			i++
		} else {
			out = append(out, ys[j])
			j++
		}
	}

	for i < len(xs) {
		out = append(out, xs[i])
		i++
	}

	for j < len(ys) {
		out = append(out, ys[j])
		j++
	}

	return out
}

func Mergesort(xs []int) []int {
	if len(xs) < 2 {
		return xs
	}
	// Split in half
	left := Mergesort(xs[:len(xs)/2])
	right := Mergesort(xs[len(xs)/2:])
	// Sort subarrays
	return sortAndMerge(left, right)
}
