package main

import "fmt"

func sqrt(n int) int {

	result := 1
	for i := 1; i < n; i++ {
		if recursivepower(result, 2) == n {
			return result
		}
		if recursivepower(result, 2) != n {
			result = result + 1
		}
	}
	return 0
}

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

	result := sqrt(64)
	fmt.Printf("Le resultat de la fonction sqrt est : %v", result)

}
