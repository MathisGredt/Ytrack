package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	var tab []bool
	for i := 0; i < len(a); i++ {
		tab = append(tab, f(a[i]))
	}
	return tab
}

func IsPrime(n int) bool {

	if n == 1 || n < 0 {
		return false
	}
	for tempo := n - 1; tempo > 1; tempo-- {
		if n%tempo == 0 {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
