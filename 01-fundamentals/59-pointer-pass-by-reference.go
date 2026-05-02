package main

import "fmt"

/**Pointer dan Operator "&"
* Adalah kemampuan membuat reference ke lokasi data dimemory yg sama, tanpa
  menduplikasi data yg sudah ada.
* Sederhananya, dg kemampuan pointer kita bisa membuat pass by reference.
* Dg menggunakan reference, kita menggunakan 1 data yg sama lalu dikirim ke semua method yg berbeda-beda. Hal ini bisa menghemat memory.
* Untuk membuat sebuah variable dg nilai pointer ke variable yg lain, 
  bisa menggunakan operator "&"" diikuti dg nama variable-nya 
**/
// Create struct
type Address struct {
	City, Province, Country 	string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1 //pointer

	address2.City  = "Bandung"

	fmt.Println("address1:", address1) // ikut berubah
	fmt.Println("address2:", address2) // berubah menjadi Bandung
}