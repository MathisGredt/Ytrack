package main

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	h := a / b
	*div = h

	j := a % 2
	*mod = j
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
