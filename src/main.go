package main

import (
	c "curso_golang_platzi/src/geometric/circle"
	"fmt"
)

type figuras2D interface {
	area() float64
}

type figuras2Dex interface {
	Area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func calcularEx(f figuras2Dex) {
	fmt.Println("Area: ", f.Area())
}

//Si los structs que tenemos en el código tienen métodos que hacen algo en común
//(Cálculos, obtener data, etc), es posible ejecutar éstos métodos usando una interfaz, de esta forma evitamos hacer código por cada struct.
func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	myCircle := c.Circle{Radio: 2}
	calcular(myCuadrado)
	calcular(myRectangulo)
	calcularEx(myCircle)

	//Reto: poner las demas operaciones con paquete y con funciones publicas y metodos privados

	//Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90} //La primera llave en blanco, la segunda con datos a instanciar
	fmt.Println(myInterface...)
}
