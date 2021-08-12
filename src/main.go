package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() { //Antes del nombre de la funcion se agrega el puntero
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	b := &a //Forma de acceder a la dirección en la memoria

	fmt.Println(a)  // Variable
	fmt.Println(b)  // Direccion de memoria
	fmt.Println(*b) // Valor que se ubica en esa direccion de memoria

	*b = 100
	fmt.Println(a) // El valor de a cambia por que se cambio el valor ubicado en esa dirección de memoria

	// Para que se utilizan los punteros: Un ejemplo es para utilizar funciones de una libreria a otro paquete de forma eficiente

	// ej:
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)
	myPC.ping()
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)
	myPC.duplicateRAM()
	fmt.Println(myPC)

	// El simbolo & accede a la memoria de la variable
	// El simbolo * accede al valor que esta guardado en memoria, usando la dirección de memoria de la variable.
}
