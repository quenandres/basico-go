package main

import (
	"fmt"
	"math"
)

func main() {
	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)

	x := 10
	y := 51

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicacion:", result)

	//Division
	result = y / x
	fmt.Println("Division:", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++
	fmt.Println("x:", x)

	// Decremental
	x--
	fmt.Println("x:", x)

	//Retos
	// -Rectangulo, trapecio y circulo
	// Area de un rectangulo
	l1ar := 5
	l2ar := 8
	result = l1ar * l2ar
	fmt.Println("Area de un rectangulo:", result)

	// Area de un trapecio
	l1ar = 5
	l2ar = 8
	ah := 5
	result = ((l1ar + l2ar) * ah) / 2
	fmt.Println("Area de un trapecio:", result)

	//Area de un circulo

	pi := math.Pi
	var rc float64 = 5.0
	var result3 = pi * math.Pow(rc, 2)
	fmt.Println("Area del circulo", result3)
}
