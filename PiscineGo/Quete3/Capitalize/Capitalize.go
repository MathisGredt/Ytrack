package main

import "fmt"

func Capitalize(sentence string) string {
	var result string
	var tempo rune
	if Islower(string(sentence[0])) == true {
		tempo = rune(sentence[0]) - 32
		result += string(tempo)
	} else {
		result += string(rune(sentence[0]))
	}
	for i := 1; i < len(sentence); i++ {
		if isprintable(string(sentence[i-1])) == false && isprintable(string(sentence[i])) == true && sentence[i] >= 97 && sentence[i] <= 122 {
			tempo = rune(sentence[i])
			tempo -= 32
			result += string(tempo)
		} else {
			tempo = rune(sentence[i])
			result += string(tempo)
		}
	}
	return result
}

func toupper(sentence string) string {
	var tempo rune
	var result string
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 97 && sentence[i] <= 122 {
			tempo = rune(sentence[i])
			tempo -= 32
			result += string(tempo)
		} else {
			result += string(rune(sentence[i]))
		}
	}
	return result
}

func isprintable(sentence string) bool {
	result := false
	for i := 0; i < len(sentence); i++ {
		if sentence[i] > 47 && sentence[i] < 58 {
			result = true
			continue
		}
		if sentence[i] > 64 && sentence[i] < 91 {
			result = true
			continue
		}
		if sentence[i] > 96 && sentence[i] < 123 {
			result = true
			continue
		}
		return false
	}
	return result
}

func Islower(sentence string) bool {
	result := false
	var sentence2 = []byte(sentence)
	for i := 0; i < len(sentence); i++ {
		result = false
		if sentence2[i] >= 97 && sentence2[i] <= 122 {
			result = true
		}
		if result == false {
			return false
		}
	}
	return result
}

func IsAlpha(sentence string) bool {
	result := false
	if len(sentence) == 0 {
		return false
	}
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 65 && sentence[i] <= 90 {
			result = true
			continue
		}
		if sentence[i] >= 97 && sentence[i] <= 122 {
			result = true
			continue
		}
		return false

	}
	return result
}

func main() {
	fmt.Println(Capitalize("Hello! how are you? how+are+you"))
}
