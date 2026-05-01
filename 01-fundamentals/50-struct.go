package main

import "fmt"

/**Struct
* Adalah template data yg digunakan untuk menggabungkan nol atau lebih tipe data lainnya dlm satu kesatuan
* Data di struct disimpan dlm field atau attribute
* Sederhananya struct adalah kumpulan dari field
* Dlm membuat field pada struct biasanya dimulai Huruf kapital istilahnya disebut PascalCase. contoh: Name, Age.
**/
type Customer struct {
	// Denifisikan field disini...
	Name, Address string
	Age 		  int
}

func main() {
	// Buar data struct -> Customer
	var sintya Customer
	sintya.Name = "Sintya Lestari"
	sintya.Address = "Bekasi"
	sintya.Age = 25

	fmt.Println(sintya)
	fmt.Println("Name Customer:", sintya.Name)
	fmt.Println("Address Customer:", sintya.Address)
	fmt.Println("Age Customer:", sintya.Age)
}