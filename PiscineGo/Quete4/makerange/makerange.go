package main

import "fmt"

func MakeRange(min, max int) []int {
	size := max - min
	var d []int
	if size > 0 {
		d := make([]int, size)
		for i := 0; i < size; i++ {
			d[i] = i + min
		}
		return d
	}
	return d
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
