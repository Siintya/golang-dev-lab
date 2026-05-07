package main

import "fmt"

/**Package fmt-https://pkg.go.dev/fmt@go1.26.2**/
func main() {
	firstName := "Sintya"
	lastName := "Lestari"

	// Println -> print line
	fmt.Println("Hello, '", firstName, lastName, "'!")

	// Printf -> print format
	fmt.Printf("Hello, '%s %s'!\n", firstName, lastName)
}
