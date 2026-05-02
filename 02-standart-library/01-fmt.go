package main

import "fmt"

/** Package fmt
https://pkg.go.dev/fmt@go1.26.2
**/
func main() {
	fmt.Println("Hello, World!")

	firstName := "Sintya"
	lastName := "Lestari"

	// Printf -> print format
	fmt.Printf("Hello '%s %s'!\n", firstName, lastName)
}