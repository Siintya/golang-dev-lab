package main

import "fmt"

// **Function Type Declarations
// Type declarations juga bisa digunakan untuk membuat alias function,
// sehingga akan menjadi mudal dlm menggunakan function sbg parameter

// Nama alias-nya Filter dg Paramater string
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "F**k Off" {
		return "***"
	} else {
		return name
	}
}

func main() {
	// Panggil Function
	sayHelloWithFilter("Sintya", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("F**k Off", filter)
}