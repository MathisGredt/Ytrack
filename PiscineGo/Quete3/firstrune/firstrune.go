package main

import "github.com/01-edu/z01"

func firstrune(sentence string) rune {
	for index, letter := range sentence {
		if index == 0 {
			return letter
		}
	}
	return 0
}

func main() {
	result := (firstrune("hello"))
	z01.PrintRune(rune(result))
}
