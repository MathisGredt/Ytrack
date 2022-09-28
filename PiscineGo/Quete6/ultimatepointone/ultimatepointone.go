package main

import "fmt"

func UltimatePointOne(n ***int) {
	h := n
	***h = 1

}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}
