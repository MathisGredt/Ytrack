package main

import "fmt"

func compare(sentence1 string, sentence2 string) int {

	if len(sentence1) > len(sentence2) {
		return -1
	}
	if len(sentence1) < len(sentence2) {
		return 1
	}
	return 0

}

func main() {
	result := compare("helloo", "hello")
	fmt.Printf("%v", result)
}
