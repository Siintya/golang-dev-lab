package main

import "fmt"

func main() {
	// **For Range
	// Bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contohnya: Array, Slice & Map

	fruits := []string{"Apple", "Watermelon", "Strawberry", "Banana"}
	for x := 0; x < len(fruits); x++ {
		fmt.Println(fruits[x])
	}

	for index, name := range fruits {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range fruits {
		fmt.Println(name)
	}

}