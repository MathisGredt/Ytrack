package main

import "fmt"

func PointOne(n *int) {
	x := n
	*x = 1

}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}
