package main

import (
	"fmt"
	"strings"
)

/**Package Strings 
Package yg berisi function-function untuk memanipulasi tipe data String
https://pkg.go.dev/strings
**/

func main() {
	// strings.Trim(string, cutset): memotong cutset di awal & akhir string
	fmt.Println(strings.Trim("   Sintya Lestari    ", " "))

	// strings.ToLower(string): membuat semua karakter string manjadi lower case
	fmt.Println(strings.ToLower("Sintya Lestari"))

	// strings.ToUpper(string): membuat semua karakter string manjadi upper case
	fmt.Println(strings.ToUpper("Sintya Lestari"))
	
	// strings.Split(string, separator): memotong string berdasarkan separator
	fmt.Println(strings.Split("Sintya Lestari", " "))

	// strings.Contains(string, search): mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains("Sintya Lestari", "tya"))

	// strings.ReplaceAll(string, from, to): mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll("Sintttyaaa", "Sintttyaaa", "Lestari"))

}



