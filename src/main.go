package main

import "fmt"

func main() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	//Println (Print + salto de linea)
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf (Imprime y agrega funcion extra al string que se le inserte)
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos) // %s para indicar String y %d para indicar int
	fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos) // %v Si no conocemos el tipo de datos

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos) // Cualquier comando generado aca, sera asignado a la variable
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("cursos: %T", cursos)
}
