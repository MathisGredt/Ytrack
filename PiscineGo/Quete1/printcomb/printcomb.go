package main

import "github.com/01-edu/z01"

func main() {
	for unité := 48; unité <= 55; unité++ {
		for dizaine := 48; dizaine <= 56; dizaine++ {
			for centaine := 48; centaine <= 57; centaine++ {
				if unité != dizaine && unité != centaine && dizaine != centaine {
					z01.PrintRune(rune(unité))
					z01.PrintRune(rune(dizaine))
					z01.PrintRune(rune(centaine))
					z01.PrintRune(rune(','))
					z01.PrintRune(rune(' '))

				}
			}

		}

	}

}
