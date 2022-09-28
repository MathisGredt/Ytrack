package main

import "fmt"

func alphacount(sentence string) int {

	result := 0
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 || sentence[i] >= 97 && sentence[i] <= 122 {
			result++
		}
	}
	return result
}

func main() {
	result := alphacount("hello")
	fmt.Println(result)
}
