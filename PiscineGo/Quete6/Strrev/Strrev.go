package main

import "fmt"

func StrRev(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
