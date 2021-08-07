package main

import "fmt"

func normalFunction(message string) { // Con un solo parametro
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) { // Reciba 3 parametros
	fmt.Println(a, b, c)
}

func returnValue(a int) int { // Que retorne una respuesta
	return a * 2
}

func doubleReturn(a int) (c, d int) { // Que retorne 2 valores
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")
	value := returnValue(5)
	fmt.Println("Value: ", value)
	value1, _ := doubleReturn(2) // Con el gúion bajo, indicamos que no nos interesa ese valor.
	fmt.Println("Value1: ", value1)
	//fmt.Println("Value2: ", value2)
}
