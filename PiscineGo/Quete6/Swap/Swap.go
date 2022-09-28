package main

import "fmt"

func Swap(a *int, b *int) {
	f := a
	g := b
	h := *f
	*f = *g
	*g = h
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
