package main

import "fmt"

// **Panic Function
// # Adalah funtion yg bisa digunakan untuk menhentikan program
// # Biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// # Saat panic function ini dipanggil, program akan terhenti. 
func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp() // tetap dieksekusi meski terjadi ERROR

	if error {
		panic("Found Error!")
	}
}

func main() {
	runApp(true)
}