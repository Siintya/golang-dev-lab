package main

import "fmt"

// **Function Return Value
// Function bisa mengembalikan data (return)
func person(name string) string {
	return "Good Morning, " + name
}

func main() {
	result := person("Sintya")
	fmt.Println(result)
}