package max

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkMax(b *testing.B) {
	numbers := make([]int, 1000000001)
	floatNumbers := make([]float64, 1000000001)
	strings := make([]string, 1000000001)
	seed := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(seed)
	for i := 0; i < b.N; i++ {
		numbers[i] = randomizer.Intn(1000)
		floatNumbers[i] = float64(randomizer.Intn(1000))
		strings[i] = strconv.Itoa(numbers[i])
	}
	b.Run("maxExplicit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxInt(numbers[i], numbers[i+1])
			maxFloat(floatNumbers[i], floatNumbers[i+1])
			maxString(strings[i], strings[i+1])
		}
	})

	b.Run("maxReflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxReflection(numbers[i], numbers[i+1])
			maxReflection(floatNumbers[i], floatNumbers[i+1])
			maxReflection(strings[i], strings[i+1])
		}
	})

	b.Run("maxTypeSwitch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxTypeSwitch(numbers[i], numbers[i+1])
			maxTypeSwitch(floatNumbers[i], floatNumbers[i+1])
			maxTypeSwitch(strings[i], strings[i+1])
		}
	})

	b.Run("maxGenerics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			maxGenerics(numbers[i], numbers[i+1])
			maxGenerics(floatNumbers[i], floatNumbers[i+1])
			maxGenerics(strings[i], strings[i+1])
		}
	})
}
