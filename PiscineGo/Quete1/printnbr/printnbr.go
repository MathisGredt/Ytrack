package main

import "github.com/01-edu/z01"

func main() {
	printnbr(200)
}

func printnbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	if number > 9 {
		printnbr(number / 10)
	}
	z01.PrintRune(rune(number%10 + 48))
}
