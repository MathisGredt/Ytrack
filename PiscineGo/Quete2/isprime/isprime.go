package main

import "fmt"

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
	result := isprime(1)
	fmt.Printf("Le resultat de la fonction isprime est : %v", result)

}
