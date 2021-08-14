package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // Definimos que termine el waitgroup de ultimas
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup // Creamos waitGroup el sync es una propiedad directa de GO

	fmt.Println("Hello")
	wg.Add(1)            // Agrega goroutine al waitgroup
	go say("World", &wg) // Le mandamos el puntero del waitgroup

	wg.Wait() // Le decimos a la waitgroup que espere mientras todas las goroutines se ejecuten

	go func(text string) { // funciones anonimas
		fmt.Println(text)
	}("Adios") //

	time.Sleep(time.Second * 1)
}
