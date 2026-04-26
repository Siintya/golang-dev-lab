package main

import "fmt"

func main() {
	// Nomor dalam array disebut index yang dimulai dari 0
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Mango"
	fruits[3] = "Kiwi"
	fruits[4] = "Strawberry"

	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
	fmt.Println(fruits[3])
	fmt.Println(fruits[4])

	// Jika ingin mengganti nilai, seperti ii
	// fruits[2] = "Grapes"
	// result akan error

	// Membuat Array Secara langsung
	var values = [3]int {20, 30, 50}
	fmt.Println(values[2])
}