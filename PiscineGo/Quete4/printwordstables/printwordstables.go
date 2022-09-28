package main

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, y := range i {
			z01.PrintRune(rune(y))
		}
		z01.PrintRune('\n')
	}
}

func SplitWhiteSpaces(sentence string) []string {
	var tab []string
	if len(sentence) != 0 {
		var tempo string
		for i := 0; i < len(sentence); i++ {
			if sentence[i] == 32 || sentence[i] == 9 || sentence[i] == 10 && tempo != "" {
				tab = append(tab, tempo)
				tempo = ""
				continue
			}
			tempo += string(sentence[i])
		}
		tab = append(tab, tempo)
	}
	return tab
}

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
