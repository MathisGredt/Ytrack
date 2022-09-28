package main

import (
	"github.com/01-edu/z01"
)

func nrune(sentence string, n int) rune {
	if n < len(sentence) {
		return rune(sentence[n])
	}
	return 48

}

func main() {
	result := (nrune("Salut", 7))
	z01.PrintRune(rune(result))
}
