package add

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkAdd(b *testing.B) {
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
	b.Run("addExplicit", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addInt(numbers[i], numbers[i+1])
			addFloat64(floatNumbers[i], floatNumbers[i+1])
			addString(strings[i], strings[i+1])
		}
	})

	b.Run("addReflection", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addReflection(numbers[i], numbers[i+1])
			addReflection(floatNumbers[i], floatNumbers[i+1])
			addReflection(strings[i], strings[i+1])
		}
	})

	b.Run("addTypeSwitch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addTypeSwitch(numbers[i], numbers[i+1])
			addTypeSwitch(floatNumbers[i], floatNumbers[i+1])
			addTypeSwitch(strings[i], strings[i+1])
		}
	})

	b.Run("addGenerics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addGenerics(numbers[i], numbers[i+1])
			addGenerics(floatNumbers[i], floatNumbers[i+1])
			addGenerics(strings[i], strings[i+1])
		}
	})

	b.Run("addGenerics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			addGenericsWithTypes(numbers[i], numbers[i+1])
			addGenericsWithTypes(floatNumbers[i], floatNumbers[i+1])
			addGenericsWithTypes(strings[i], strings[i+1])
		}
	})
}
