package main

import "github.com/01-edu/z01"

func main() {
	var virgule bool = true
	for dizaine1 := 48; dizaine1 < 58; dizaine1++ {
		for unite1 := 48; unite1 < 58; unite1++ {
			for dizaine2 := 48; dizaine2 < 58; dizaine2++ {
				for unite2 := 48; unite2 < 58; unite2++ {
					if dizaine1*10+unite1 < dizaine2*10+unite2 {
						if !virgule {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}
						z01.PrintRune(rune(dizaine1))
						z01.PrintRune(rune(unite1))
						z01.PrintRune(rune(' '))
						z01.PrintRune(rune(dizaine2))
						z01.PrintRune(rune(unite2))
						virgule = false
					}
				}
			}
		}
	}
}
