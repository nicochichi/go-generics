package max

import "golang.org/x/exp/constraints"

func maxGenerics[T int | float64 | string | constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}
