package main

import "fmt"

// **Function Parameter
// Bisa menambahkan lebih dari satu parameter di function
// Parameter TIDAK WAJIB, namun jika sudah menambhkan parameter di function
// Maka ketika memanggil function tsb, harus memasukkan data ke parameternya.
func person(firstName string, lastName string, age int) {
	fmt.Println("Name:", firstName, lastName, "\nAge:", age)
}

func main() {
	person("Sintya", "Lestari", 25)
}