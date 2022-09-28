package main

import "fmt"

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == -1 {
			result = false
		}
	}
	return result
}

func CompareNum(a int, b int) int {
	result := 0
	if a > b {
		result = -1
	}
	if a < b {
		result = 1
	}
	return result
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(CompareNum, a1)
	result2 := IsSorted(CompareNum, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
