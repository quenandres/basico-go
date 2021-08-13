package main

import "fmt"

func main() {
	//Concurrencia:
	// La concurrencia lidia con multiples cosas al mismo tiempo, mientras que el paralelismo hace multiples cosas al mismo tiempo.
	m := make(map[string]int)

	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	m["5"] = 5
	m["6"] = 6
	m["7"] = 7
	m["8"] = 8
	m["9"] = 9
	m["10"] = 10

	for _, value := range m {
		fmt.Printf("%d\n", value)
	}
}
