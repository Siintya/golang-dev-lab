package main

import "fmt"

func main() {
	// **Else If Expression (Percabangan)
	nilai := 55

	if nilai > 80 {
		// Blok if akan dieksekusi ketika kondisi if bernilai true 
		fmt.Println("Selamat Kamu LULUS!")
	} else if nilai == 75 {
		fmt.Println("Kamu Remidial, Semangat!")
	} else {
		// Blok else akan dieksekusi jika bernilai False
		fmt.Println("Maaf, Kamu TIDAK LULUS!")
	}
}