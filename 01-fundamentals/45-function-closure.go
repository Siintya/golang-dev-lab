package main

import "fmt"

// **Closure Function
// Yaitu sebuah fitur di mana sebuah fungsi "menangkap" (capture) dan 
// mereferensikan variabel-variabel yg berada di luar lingkup (scope) tubuh fungsinya sendiri.
func main() {
	counter := 0

	// Buat variable dg nama increment yg value-nya sebuah fungsi anonim ini adalah closure
	increment := func() {
		fmt.Println("Increment", counter)
		counter++ // Ia "menangkap" variabel counter dari scope luar
	}

	// Panggil
	increment()
	increment()
	increment()
	increment()
	fmt.Println(counter)
}