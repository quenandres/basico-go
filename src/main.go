package main

import (
	pk "curso_golang_platzi/src/mypackage" //Alias a la izquierda con un espacio.
	"fmt"
)

func main() {
	// Se indica si una variable, funcion o tipo de datos puede tener acceso publico o privado
	// Si la primera letra esta en mayuscula significa que es publico, de lo contrario es privado.
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	//Go nos indica que no podemos acceder a esa variable por no ser metodo publico.
	fmt.Println(myCar)

	//var myOtherCar pk.CarPrivate
	pk.PrintMessage("Hola Platzi")
	pk.printMessage("Hola")
}
