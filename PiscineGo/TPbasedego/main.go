package main

import "fmt"

func CheckAge(age int) bool {

	if age >= 18 {
		return true
	}
	return false
}

func CheckIdentity(age int, sexe string) string {

	result := ""
	if CheckAge(age) == true && sexe == "femme" {
		result = "Tu peux entrer, c’est 10 €"
	}
	if CheckAge(age) == true && sexe != "femme" {
		result = "Tu peux entrer, c’est 15 €"
	}
	if CheckAge(age) == false {
		result = "Sortez !"
	}
	return result
}

func main() {
	fmt.Println(CheckAge(17))
	fmt.Println(CheckAge(19))
	fmt.Println(CheckIdentity(21, "femme"))
	fmt.Println(CheckIdentity(22, "homme"))
	fmt.Println(CheckIdentity(17, "femme"))
}
