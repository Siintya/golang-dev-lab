package main

import "fmt"

// **Anonymous Function (Function Tanpa Nama)
// Berbeda dengan fungsi biasa yang dideklarasikan menggunakan kata kunci func diikuti nama fungsinya, anonymous function biasanya dibuat tepat di tempat ia akan digunakan. 
// Ini sangat berguna untuk logika singkat yg tdk perlu dipanggil berulang kali di berbagai tempat.
// Didefinisikan menggunakan func() { ... }.

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// 1. buat varible blacklist dg value-nya adalah function anonymous
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Sintya", blacklist)

	// 2. buat function anonymous scr langsung sbg parameter
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}