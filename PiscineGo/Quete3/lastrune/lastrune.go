package main

import (
	"github.com/01-edu/z01"
)

func lastrune(sentence string) rune {

	return rune(sentence[len(sentence)-1])
}

func main() {
	result := (lastrune("Salut"))
	z01.PrintRune(rune(result))
}
