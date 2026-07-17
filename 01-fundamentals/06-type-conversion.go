package main

import "fmt"

func main() {
	var angkaBulat int = 568

	var angkaDesimal float64 = float64(angkaBulat)

	fmt.Println("Angka bulat:", angkaBulat)
	fmt.Printf("Angka desimal: %.1f\n", angkaDesimal)
}
