package main

import (
	"fmt"
)

func isprintable(sentence string) bool {
	result := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] > 47 && sentence[i] < 58 {
			result = true
			continue
		}
		if sentence[i] > 64 && sentence[i] < 91 {
			result = true
			continue
		}
		if sentence[i] > 96 && sentence[i] < 123 {
			result = true
			continue
		}
		return false
	}
	return result
}

func main() {
	fmt.Println(isprintable("151GRBj"))

}
