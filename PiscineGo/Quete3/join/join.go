package main

import "fmt"

func join(elems []string, sep string) string {

	result := ""
	counter := len(elems) - 1
	for i := 0; i < len(elems); i++ {
		result += string(elems[i])
		if counter > 0 {
			result += ":"
		}
		counter--
	}
	return result
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(join(elems, ":"))

}
