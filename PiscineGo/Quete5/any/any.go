package main

import "fmt"

func Any(f func(string) bool, a []string) bool {
	result := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			result = true
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
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
