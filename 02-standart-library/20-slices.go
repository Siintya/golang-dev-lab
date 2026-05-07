//go:build ignore

package main

import (
	"fmt"
	"slices"
)

/**Package Slices
* Di golang terbaru, terdapat fitur Generic Programming
* Fitur ini membuat kita bisa membuat parameter dg tipe yg bisa berubah-ubah,
  tanpa harus menggunakan interface kosong / any
* Salah satu package yg menggunakan fitur ini yaitu pakcage slices
* package ini digunkan untuk memanipulasi data di slice
https://pkg.go.dev/slices
**/

func main() {
	names := []string{"Jhon", "Doe", "Jane", "Paul"}
	values := []int{100, 95, 80, 90}

	fmt.Println("Min:", slices.Min(names))
	fmt.Println("Min:", slices.Min(values))

	fmt.Println("Max:", slices.Max(names))
	fmt.Println("Max:", slices.Max(values))

	fmt.Println("Contains:", slices.Contains(names, "Sintya"))

	fmt.Println("Index:", slices.Index(names, "Sintya"))
	fmt.Println("Index:", slices.Index(names, "Jane"))

}
