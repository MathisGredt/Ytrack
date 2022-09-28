package main

import "fmt"

func interativefactorial(index int) int {

	result := 1
	for i := 1; i <= index; i++ {
		result = result * i
	}
	return result
}

func main() {

	result := interativefactorial(4)
	fmt.Printf("Le resultat de la fonction interactivefactorial est : %v", result)

}
