package main

import "fmt"

func findnextprime(n int) int {

	for i := n; i > 0; i++ {
		if isprime(i) {
			return i
		}
	}
	return 0

}

func isprime(n int) bool {

	if n == 1 || n < 0 {
		return false
	}
	for tempo := n - 1; tempo > 1; tempo-- {
		if n%tempo == 0 {
			return false
		}
	}
	return true
}

func main() {
	result := findnextprime(7)
	fmt.Printf("Le resultat de la fonction findnextprime est : %v", result)

}
