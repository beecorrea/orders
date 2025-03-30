package order

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
