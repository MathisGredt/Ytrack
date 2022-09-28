package main

import "fmt"

func tolower(sentence string) string {
	var tempo rune
	var result string
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 {
			tempo = rune(sentence[i])
			tempo += 32
			result += string(tempo)
		} else {
			result += string(rune(sentence[i]))
		}
	}
	return result
}

func main() {
	fmt.Println(tolower("aGHju"))
}
