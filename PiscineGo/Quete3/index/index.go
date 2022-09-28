package main

import "fmt"

func Index(sentence string, tofind string) int {
	result := 0
	counter := 0
	var index int
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == tofind[result] {
			result += 1
			counter += 1
			if counter == 1 {
				index = i
			}
		} else {
			result = 0
		}
		if result == 0 {
			counter = 0
		}

		if result == len(tofind) {
			return index
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola", "hol"))

}
