package main

import "fmt"

func main() {
	var fruit1 = "Apple"
	var fruit2 = "Apple"

	var result bool = fruit1 == fruit2
	fmt.Println(result)

	var result2 bool = fruit1 != fruit2
	fmt.Println(result2)

	// >
	var x = 5
	var y = 10

	var result3 bool = x > y
	fmt.Println(result3)

	// <
	var result4 bool = x < y
	fmt.Println(result4)
}