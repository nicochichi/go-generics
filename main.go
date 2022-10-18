package main

import (
	"fmt"

	"nicochichi/generics/collections"
	"nicochichi/generics/convert"
)

func main() {
	fmt.Println(convert.ToString(uint64(234)))

	m := map[string]int64{"foo": 1, "bar": 2, "test": 3, "anotherValue": 4}

	fmt.Println(collections.Keys(m))

	test := []string{"test", "foo", "bar"}

	fmt.Println(collections.Contains(test, "foo"))

	fmt.Println(collections.Contains(test, "HELLO"))
}
