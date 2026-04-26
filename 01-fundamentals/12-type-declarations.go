package main

import "fmt"

func main() {
	// Type Declarations
	// NoKTP adalah Alias dari tipe data string
	type NoKTP string

	var ktpSintya NoKTP = "2222222"
	fmt.Println(ktpSintya)
	fmt.Println(NoKTP("66666666"))
	

}