package main

import "fmt"

func main() {
	// **Switch Expression
	// Biasanya digunakan untuk melakukan pengecekan ke kondisi dalam satu variable
	name := "Ajeng"

	switch name {
	case "Sintya":
		// Blok ini akan dieksekusi jika variable bernilai true
		fmt.Println("Hello, Sintya.")
	case "Lestari":
		// Blok ini akan dieksekusi jika variable bernilai true
		fmt.Println("Hello, Lestari.")
	default:
		// Blok ini akan dieksekusi jika variable bernilai false
		fmt.Println("Hai, Boleh Kenalan?")
	}

	// **Switch dg Short Statement
	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// ** Switch tanpa kondisi
	// Switch dg kondisi seperti diatas ini TIDAK WAJIB
	// Jika tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tsb di setiap case-nya

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default: 
		fmt.Println("Nama Sudah Benar")
	}
}