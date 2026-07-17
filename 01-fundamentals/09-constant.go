package main

import "fmt"

func main() {
	// Constant: variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai.
	const firstName string = "Sintya"

	// ERROR
	firstName = "Budi"
	fmt.Println(firstName)

}
