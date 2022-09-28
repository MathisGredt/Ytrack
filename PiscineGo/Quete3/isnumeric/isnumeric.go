package main

import (
	"fmt"
)

func Isnumeric(sentence string) bool {

	for i := 0; i < len(sentence); i++ {
		if sentence[i] < 48 || sentence[i] > 57 {
			return false
		}
	}
	return true
}

func main() {
	result := Isnumeric("89451gg8")
	fmt.Println(result)

}
