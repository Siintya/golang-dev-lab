package main

import "fmt"

/** Pointer di Function
* Saat membuat parameter di function, scr default adalah pass by value, 
  artinya data akan di copy lalu dikirim ke function tsb. Oleh karena itu,
  jika mengubah data di dlm function tsb, data yg asli-nya tdk pernah berubah.
* Hal ini membuat variable menjadi aman, karna tdk akan bisa diubah
* Namun, terkadang kita ingin membuat function yg bisa mengubah
  data asli parameter tsb dengan cara menggunakan pointer di function.
* Untuk menjadikan sebuah parameter sbg pinter, kita bisa menggunakan
  operator * di parameter-nya.
**/
type Address struct {
	City, Province, Country	string
}

func ChangeAddressCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address *Address = &Address{"Bekasi", "Jawa Barat", ""}
	ChangeAddressCountry(address)

	fmt.Println(address)
}