package add

type types interface {
	int | float64 | string
}

func addGenericsWithTypes[T types](a, b T) T {
	return a + b
}

func addGenerics[T int | float64 | string](a, b T) T {
	return a + b
}
