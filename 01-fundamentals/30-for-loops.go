package main

import "fmt"

func main() {
	// **For Loops (Perulangan)
	// for condition { blok code yg akan terus dieksekusi jika kondisi bernilai true}

	// counter := 10

	// for counter <= 20 {
	// 	fmt.Println("Perulangan ke",counter)
	// 	counter++
	// }

	// **For dengan Statement
	// Terdapat 2 statement:
	// 1. init statement, yaitu statement sebelum for dieksekusi
	// 2. post statement, yaitu statement yg akan selalu dieksekusi di akhir tiap perulangan
	// for init statement; condition; post statement { blok code }

	for counter := 5; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}


}