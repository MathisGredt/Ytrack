package main

import (
	"fmt"
)

func CountIf(f func(string) bool, tab []string) int {
	result := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			result += 1
		}
	}
	return result
}

func IsNumeric(sentence string) bool {
	for i := 0; i < len(sentence); i++ {
		if sentence[i] < 48 || sentence[i] > 57 {
			return false
		}
	}
	return true
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
