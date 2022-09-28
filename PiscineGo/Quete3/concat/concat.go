package main

import "fmt"

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	result := Concat("Hello! ", "How are you?")
	fmt.Println(result)

}
