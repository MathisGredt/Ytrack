package main

import (
	"fmt"
	"os"
)

func main() {

	for i := 1; i < len(os.Args); i++ {
		for j := 1; j < len(os.Args); j++ {
			if os.Args[i] < os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}
	fmt.Println(os.Args)
}
