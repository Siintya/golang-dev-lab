package main

import "fmt"

// **Function
// Adalah sebuah blok kode yg sengaja dibuat dlm program agar bisa dieksekusi/digunakan scr berulang-ulang
// Cara membuat function : menggunakan keyword func lalu diikuti dg nama function-nya & blok kode isi function-nya
// func nameFunction() { blok code }
func sayHello() {
	fmt.Println("Hello, Guys!!")
}

func main() {
	// Panggil function
	sayHello()
}