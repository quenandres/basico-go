package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// And
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// Or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Alguna de las 2 es verdadera")
	}

	//Reto con funcion definir si un numero es par o impar
	reto1 := parImpar(4)
	if reto1 == true {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
	//Reto 2 login de usuario y contraseña
	reto2 := login("Carlos", "1223456")
	if reto2 == true {
		fmt.Println("Login :3")
	} else {
		fmt.Println("Vuelva a intentar")
	}
}

func parImpar(n int) bool {
	return n%2 == 0
}

func login(name string, password string) bool {
	if name == "Carlos" && password == "123456" {
		return true
	}
	return false
}
