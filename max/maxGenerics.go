package max

func maxGenerics[T int | float64 | string](a, b T) T {
	if a > b {
		return a
	}

	return b
}
