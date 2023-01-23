package main

import (
	"fmt"

	"nicochichi/generics/collections"
	"nicochichi/generics/convert"
	"nicochichi/generics/domain"
)

func main() {
	// Convertimos un valor uint64 a string usando el metodo generico que creamos.
	// Podemos usar el comodin ~ para incluir a todos los tipos que tienen como base el tipo que definimos
	fmt.Println("Valor uint64 convertido a string: ", convert.ToString(uint64(234)))

	// Podemos remover el tipo shipmentID cuando llamamos a la funcion porque el compilador puede inferir el tipo segun el parametro.
	fmt.Println("Valor ShipmentID convertido a string: ", convert.ToString[domain.ShipmentID](domain.ShipmentID(234)))

	// Print  de todos los keys del map.
	m := map[string]int64{"foo": 1, "bar": 2, "test": 3, "anotherValue": 4}
	fmt.Println("Keys del map: ", collections.Keys(m))

	// Busca si el slice contiene el valor ingresado.
	test := []string{"test", "foo", "bar"}
	fmt.Println("La coleccion contiene el valor 'foo': ", collections.Contains[string](test, "foo"))
	fmt.Println("la coleccion contiene el valor 'HELLO':", collections.Contains(test, "HELLO"))

	// Uso de generics en structs.
	user1 := domain.UserGeneric[domain.DataStructure]{
		ID:   123,
		Name: "test",
		Data: domain.DataStructure{Age: 12, Height: 12.2},
	}

	fmt.Println("User 1 usando data structure: ", user1)

	user2 := domain.UserGeneric[int64]{
		ID:   123,
		Name: "test",
		Data: 1,
	}
	fmt.Println("User 2 usando int64", user2)

	// Uso de Generics para definir tipos de maps.
	newMap := make(collections.CustomMap[string, int])
	newMap["hello"] = 2
	newMap["world"] = 3

	fmt.Println("custom map: ", newMap)
	fmt.Println("Keys del custom map: ", collections.Keys(newMap))

	// Implementacion de una linked list
	list := collections.List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)
	list.Push(2)
	fmt.Println("contenido de la lista: ", list.GetAll())
}
