package mypackage

import "fmt"

// Se indica si una variable, funcion o tipo de datos puede tener acceso publico o privado
// Si la primera letra esta en mayuscula significa que es publico, de lo contrario es privado.
// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir un mensaje
func PrintMessage(msj string) {
	fmt.Println("Hola")
	fmt.Println(msj)
}

func printMessage(msg string) {

}
