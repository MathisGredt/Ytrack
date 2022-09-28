package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	j := a
	k := b
	x := *j / *k
	*k = *j % *k
	*j = x

}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
