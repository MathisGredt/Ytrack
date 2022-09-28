package main

import "fmt"

func Split(sentence string, sep string) []string {
	var result []string
	var tempo string
	var result1 string
	for i := 0; i < len(sentence); i++ {
		tempo = string(sentence[i])
		result1 += tempo
		if sentence[i] == 72 && sentence[i+1] == 65 {
			result = append(result, result1)
		} else {
			tempo = string(sentence[i])
			result1 += tempo
		}
	}
	return result

}

func main() {
	s := "HelloHAHowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
