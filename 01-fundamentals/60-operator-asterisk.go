package main

import "fmt"

/**Asterick Operator atau Operator "*"
* saat mengubah variable ponter, maka yg berubah hanya variable tsb.
* semua variable yg mengacu ke data yg sama tdk akan berubah
* Jika ingin mengubah-nya, bisa menggunakan operator "*"
**/
type Address struct {
	City, Province, Country	string
}

func main() {
	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	// Ubah nilai City
	address2.City = "Bandung"

	// pointer
	address2  = &Address{"Yogyakarta", "Jawa Tengah", "Indonesia"}
	// operator *
	*address3 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address1) // berubah
	fmt.Println(address2) // tdk akan berubah
	fmt.Println(address3) // berubah, karna address3 := &address1
}