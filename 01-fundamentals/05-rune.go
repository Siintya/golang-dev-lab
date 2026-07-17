package main

import "fmt"

func main() {
	str := "Go😊"
	runes := []rune(str)

	fmt.Println("Jumlah karakter sebenarnya:", len(runes)) // Output: 3
	fmt.Printf("Karakter ketiga: %c\n", runes[2])          // Output: 😊
}
