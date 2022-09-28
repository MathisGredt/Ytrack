package main

import "fmt"

func BasicAtoi2(s string) int {
	return TrimAtoi(s)
}

func TrimAtoi(sentence string) int {
	sign := 1
	var result int
	var i int
	premier_chiffre := true
	for i < len(sentence) {
		if sentence[i] == '-' && premier_chiffre == true {
			sign = -1
		}
		if sentence[i] == 32 {
			return 0
		} else if sentence[i] > 48 && sentence[i] <= 57 {
			result = result*10 + int(sentence[i]-48)
			premier_chiffre = false
		}
		i++
	}
	return result * sign
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
