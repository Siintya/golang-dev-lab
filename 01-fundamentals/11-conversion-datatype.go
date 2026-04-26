package main

import "fmt"

func main() {
	// Konversi Tipe Data Number
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	// result: -32768
	// number over flow, jadi jika melebihi kapasitas dia akan balik lagi ke belakang.

	// Konversi Tipe Data String
	var name = "Sintya Lestari"
	var e uint8 = name[5]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}