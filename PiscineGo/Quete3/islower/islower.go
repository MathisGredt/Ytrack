package main

import (
	"fmt"
)

func Islower(sentence string) bool {
	result := false
	var sentence2 = []byte(sentence)
	for i := 0; i < len(sentence); i++ {
		result = false
		if sentence2[i] >= 97 && sentence2[i] <= 122 {
			result = true
		}
		if result == false {
			return false
		}
	}
	return result
}

func main() {
	result := Islower("f")
	fmt.Println(result)

}
