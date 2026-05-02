package main

import "fmt"

/** Operator new
* function ini bisa digunakan untuk membuat pointer
* function new ini hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
**/
type Address struct {
	City, Province, Country	string
}

func main() {
	var alamat1 = new(Address)
	var alamat2 = alamat1

	alamat2.Country = "Singapura"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}