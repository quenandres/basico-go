package mypackage

import "fmt"

// Se indica si una variable, funcion o tipo de datos puede tener acceso publico o privado
// Si la primera letra esta en mayuscula significa que es publico, de lo contrario es privado.
// CarPublic Car con acceso publico
type MyPCPublic struct {
	Ram   int
	Disk  int
	Brand string
}

func (MyPC MyPCPublic) SeePC() {
	fmt.Printf("El equipo tiene %d ram %d disco duro y es de marca %v", MyPC.Ram, MyPC.Disk, MyPC.Brand)
}
