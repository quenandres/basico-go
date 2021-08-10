package main

import "fmt"

func main() {
	// Array: Los arrays son inmutables por eso son estaticos
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array)) // len brinda el tamaño del array y cap la capacidad maxima del array

	// Slice: Muy similar al array pero no se indica la capacidad
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])  // El numero es exclusivo
	fmt.Println(slice[2:4]) // El primero es inclusivo y el segundo es exclusivo
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Los tres puntos indican que va a descomponer la lista nueva y agregara los elementos a el slice
	fmt.Println(slice)
}
