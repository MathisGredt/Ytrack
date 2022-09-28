package main

import "fmt"

func recursivefactorial(index int) int {

	if index == 1 {
		return 1
	}
	if index > 0 {
		return index * recursivefactorial(index-1)
	}
	return 0
}

func main() {

	result := recursivefactorial(4)
	fmt.Printf("Le resultat de la fonction recursivefactorial est : %v", result)

}
