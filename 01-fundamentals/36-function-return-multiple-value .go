package main

import "fmt"

// **Returning Multiple Value
//  untuk mengembalikan multiple value, harus menulis semua tipe data return value-nya di function tsb.
func getPerson() (string, string, int)  {
	return "Sintya", "Lestari", 25
}

func main() {
	firstName, lastName, age := getPerson()
	fmt.Println("Fullname:", firstName, lastName, "\nAge:", age)
}