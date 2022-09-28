package main

import "fmt"

func TrimAtoi(sentence string) int {
	sign := 1
	var result int
	var i int
	premier_chiffre := true
	for i < len(sentence) {
		if sentence[i] == '-' && premier_chiffre == true {
			sign = -1
		} else if sentence[i] > 48 && sentence[i] <= 57 {
			result = result*10 + int(sentence[i]-48)
			premier_chiffre = false
		}
		i++
	}
	return result * sign
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
