package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string { // Funcion que tiene definido el struct, define la funcion String y se le dice que retornara un string, Sprintf crea un string para ser impreso.
	return fmt.Sprintf("Tengo %d GB RAM %d GB de disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {

	myPC := pc{ram: 16, disk: 100, brand: "msi"}
	fmt.Println(myPC)

}
