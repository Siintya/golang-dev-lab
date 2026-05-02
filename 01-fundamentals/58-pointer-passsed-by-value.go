package main

import "fmt"

/**Pointer - Passed by Value
Secara Default di Golang semua variable itu di passing by value, bukan by reference.
Artinya, Jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
sebenarnya yg dikirim adalah duplikasi value-nya. Jadi  diduplicate dulu baru dikirim data.
**/
// Create struct
type Address struct {
	City, Province, Country 	string
}
func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := address1
	address3 := address1

	address2.City    = "Bandung"
	address3.Country = "Malaysia"

	fmt.Println(address1) // address 1 tdk berubah
	fmt.Println(address2)
	fmt.Println(address3)
}