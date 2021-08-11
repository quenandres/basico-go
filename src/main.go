package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	// En Go no existe la estructura de clase en su lugar se usan los struct
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra forma de instanciar
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
