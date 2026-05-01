package main

import "fmt"

/**Recover Function
# Adalah function yg bisa digunakan untuk menangkap data panic
# Dg recover proses panic akan terhenti, sehingga program akan tetap berjalan.
**/

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Found ERROR!")
	}
}

func main() {
	runApp(true)
	fmt.Println("Sintya")
}