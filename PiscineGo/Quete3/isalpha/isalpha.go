package main

import (
	"fmt"
)

func IsAlpha(sentence string) bool {
	result := false
	if len(sentence) == 0 {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 {
			result = true
			continue
		}
		if sentence[i] >= 97 && sentence[i] <= 122 {
			result = true
			continue
		}
		return false

	}
	return result
}

func main() {
	result := IsAlpha("kjb")
	fmt.Println(result)

}
