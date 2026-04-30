package main

import "fmt"

// **Function as Parameter
// filter func(string) string -> function dengan 1 buah parameter string & return-nya string.
func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
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