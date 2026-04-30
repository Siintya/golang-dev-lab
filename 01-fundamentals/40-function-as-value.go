package main

import "fmt"

// **Function Value
// Di Go, Function adalah first class citizen, artinya fungsi dianggap sebagai tipe data.
// Jad bisa disimpan pada sebuah variable atau dijadikan sebuah value (nilai) 
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Membuat function sbg value pada variable goodbye
	goodbye := getGoodBye
	example := getGoodBye
	fmt.Println(goodbye("Sintya"))
	fmt.Println(example("Lestari"))
}