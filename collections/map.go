package collections

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type CustomMap[T comparable, V int | string] map[T]V

func (c CustomMap[T, V]) Keys() []T {
	keys := make([]T, 0, len(c))
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}
