package main

import (
	"fmt"
)

func IsUpper(sentence string) bool {
	result := false
	var sentence2 = []byte(sentence)
	for i := 0; i < len(sentence); i++ {
		result = false
		if sentence2[i] >= 65 && sentence2[i] <= 90 {
			result = true
		}
		if result == false {
			return false
		}
	}
	return result
}

func main() {
	result := IsUpper("GHPG!")
	fmt.Println(result)

}
