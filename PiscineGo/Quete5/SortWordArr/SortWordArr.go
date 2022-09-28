package main

import "fmt"

func SortWordArr(a []string) {
	if len(a) > 1 {
		for i := 0; i < len(a)-1; i++ {
			for j := 0; j < len(a)-1; j++ {
				if Compare(a[j], a[j+1]) == 1 {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
	}
}

func Compare(str1, str2 string) int {
	min_len := GetMinLen(len(str1), len(str2))
	for i := 0; i < min_len; i++ {
		if str1[i] > str2[i] {
			return 1
		} else if str1[i] < str2[i] {
			return -1
		}
	}
	if len(str1) > len(str2) {
		return 1
	} else if len(str1) < len(str2) {
		return -1
	}
	return 0
}

func GetMinLen(len1, len2 int) int {
	if len1 > len2 {
		return len2
	}
	return len1
}

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)
	fmt.Println(result)
}
