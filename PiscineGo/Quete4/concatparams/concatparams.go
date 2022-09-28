package main

import "fmt"

func ConcatParams(args []string) string {
	var result string
	for i := 0; i < len(args); i++ {
		result += string(args[i]) + "\n"
	}
	return result
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
