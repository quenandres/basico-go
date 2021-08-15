package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message 1"
	c <- "Message 2"
	fmt.Println(len(c), cap(c)) // len = Tamaño de las goroutines en el canal. cap = capacidad del canal

	// Range y close
	close(c) // Le dice a routime que va a cerrar el canal, diciendo que el canal no va a recibir mas información
	// Lo mejor es cerrarlo si sabemos que lo se utilizara mas
	//c <- "Message 3"

	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string) // No se especifica la cantidad maxima que tendra
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
