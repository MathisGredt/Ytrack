package main

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

func PrintNbr(number int) {
	if number < 0 {
		z01.PrintRune('-')
		number = -number
	}
	if number > 9 {
		PrintNbr(number / 10)
	}
	z01.PrintRune(rune(number%10 + 48))
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
