package main

import "fmt"

func main() {
	// Augmented Assignments
	var a = 10

	a += 30
	fmt.Println(a)

	a -= 5
	fmt.Println(a)
	
	a *= 40
	fmt.Println(a)

	a /= 2
	fmt.Println(a)

	a %= 3
	fmt.Println(a)
}