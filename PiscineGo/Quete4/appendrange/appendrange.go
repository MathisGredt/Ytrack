package main

import "fmt"

func AppendRange(min, max int) []int {
	var tableau []int
	var tableau2 []int
	for i := min; i < max; i++ {
		if min < max {
			tableau = append(tableau, i)
		}
		if min > max {
			return tableau2
		}

	}
	return tableau
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
