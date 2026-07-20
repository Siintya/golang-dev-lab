package main

import "fmt"

func main() {
	/**Map adalah tipe data lain yg berisi kumpulan data yg sama, namun bisa menentukan tipe data index yg akan digunakan
	Membuat map: map[typeData]typeData{"key": "value",}
	**/

	// Cara 1 membuat map
	// var person := map[string]string{}
	// person["name"] = "Sintya"
	// person["address"] = "Bekasi"

	// Cara 2 membuat Map
	person := map[string]string{
		"name":    "Sintya",
		"address": "Bekasi",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["usia"]) // result kosong
	fmt.Println(person)

	value, ok := person["usia"]

	if ok {
		fmt.Println("Key ketemu! Nilainya adalah:", value)
	} else {
		fmt.Println("Key tidak ditemukan di dalam map.")
	}

	_, exists := person["usia"]
	if exists {
		fmt.Println("Key ketemu! Nilainya adalah:", value)
	} else {
		fmt.Println("Key tidak ditemukan di dalam map.")
	}

}
