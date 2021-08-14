package main

import (
	"fmt"
)

func say(text string, c chan<- string) { // Con el simbolo <- indicamos que este canal solo sera de entrada de datos para salida se debe colocar de lado izq, <-chan
	c <- text
}

func main() {
	c := make(chan string, 1) // Se crea el canal
	fmt.Println("Hello")
	go say("Bye", c)
	fmt.Println(<-c)
}
