package main

import "fmt"

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
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))

}
