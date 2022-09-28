package main

import "fmt"

func Basicjoin(elems []string) string {

	result := ""
	for i := 0; i < len(elems); i++ {
		result += string(elems[i])
	}
	return result
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Basicjoin(elems))

}
