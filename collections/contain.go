package collections

func Contains[T comparable](col []T, valueToFind T) bool {
	for _, value := range col {
		if value == valueToFind {
			return true
		}
	}
	return false
}
