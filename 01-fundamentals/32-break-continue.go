package main

import "fmt"

func main() {
	// **Break & Continue
	// Break digunakan ntuk menghentikan seluruh perulangan
	// Continue digunakan untuk menghentikan perulangan yg berjalan-
	// dan langsung melanjutkan ke perulangan selanjutnya

	// Break
	for x := 0; x < 10; x++ {
		if x == 5 {
			fmt.Println("Angka telah ditemukan ", x)
			break
		}

		fmt.Println("Perulangan ke", x)
	}

	// Continue
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			fmt.Println("Angka telah ditemukan ", x)
			continue
		}

		fmt.Println("Perulangan ke", x)
	}

}