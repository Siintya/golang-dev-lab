package main

import "fmt"

func main() {
	// **If dg Short Statement
	// Yg dimanakan short statement ini -> if var2 := len(var1); var2 > 5 
	// Cocok untuk memebuat statement yang sederhana sebelum melakukan pengecekan kondisi
	name := "Sintya Lestari"

	if length := len(name); length > 10 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}