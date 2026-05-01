package main

import "fmt"

/**Type Assertions menggunakan Switch
* Type Assertion bersifat "keras" (strict). Tipe data yang diminta harus sama persis dengan tipe data asli yang dimasukkan ke dalam interface kosong tsb.
* Jika panic & tidak ter-recover, maka otomatis program kita akan mati
* Agar lebih aman, sebaiknya menggunakan switch expression untuk melakukan type assertions
**/

// Interface kosong
func random() any {
	// return 89
	return "Hello World!"
}

func main() {
	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Unknown")
	}
}