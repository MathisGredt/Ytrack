package main

import "fmt"

func recursivepower(nb int, power int) int {

	if power == 1 {
		return nb
	}
	if power > 1 {
		return nb * recursivepower(nb, power-1)
	}
	return 0
}

func main() {

	result := recursivepower(8, 3)
	fmt.Printf("Le resultat de la fonction recursivepower est : %v", result)
}
