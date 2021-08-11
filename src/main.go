package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice { //Escapamos el indice con el _
		fmt.Println(valor)
	}

	str := strings.ToLower("amA")
	isPalindromo(str)

}
