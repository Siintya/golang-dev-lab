package main

import "fmt"

/**Type Assertions
* Yaitu sebuah fitur yg memiliki kemampuan merubah tipe data menjadi tipe data yg diinginkan
* Fitur ini sering kali digunakan ketika kita bertemu dg data interface kosong
**/

// Interface kosong
func random() any {
	return "OK!"
}

func main() {
	var result any = random()

	// Type Assetions -> harus sama dg nilai tipe data yang dikembalikan pada interface kosong tsb
	var resultString string = result.(string)
	fmt.Println(resultString)

	// Hasil ERROR panic karna pada interface kosong random() mengembalikan nilai yg bertipe data string bukan integer
	var resultIn int = result.(int) // panic
	fmt.Println(resultIn)
}