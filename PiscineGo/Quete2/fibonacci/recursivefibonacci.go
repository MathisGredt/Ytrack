package main

import "fmt"

func recursivefibonacci(n int) int {

	if n < 0 {
		return -1
	}
	if n == 1 {
		return 1
	}
	if n > 1 {
		return recursivefibonacci(n-1) + recursivefibonacci(n-2)
	}
	return 0
}

func main() {
	result := recursivefibonacci(6)
	fmt.Printf("Le resultat de la fonction recursivefibonacci est : %v", result)
}
