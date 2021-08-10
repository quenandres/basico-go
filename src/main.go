package main

import "fmt"

func main() {

	// Defer: Se podria usar como defer para cerrar la conexion con una bd, cerrar archivo que ya se haya leido.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue: Se utiliza cuando una condicion dada, para que continue a pesar de un error
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		// Break: Cuando se quiere detener la ejecución del codigo
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
