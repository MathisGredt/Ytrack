package main

import "fmt"

func iterativepower(nb int, power int) int {

	result := 1
	for i := power; i > 0; i-- {
		result *= nb
	}
	return result

}

func main() {

	result := iterativepower(4, 3)
	fmt.Printf("Le resultat de la fonction iterativepower est : %v", result)

}
