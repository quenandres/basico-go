package main

import "fmt"

func main() {
	//Los map utilizan concurrencia de forma nativa por lo que son mas eficientes que los slices
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
