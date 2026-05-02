package main

import (
	"01-fundamentals/helper"
	"fmt"
)

/**Package & Import
* Package adalah hanya directory folder dari kode program yg kita buat
* scr standar, file go hanya bisa mengakses file go lainnya yg berada dlm
  package yg sama.
* Jika ingin mengakses file go yg berada diluar package, maka gunakan import
**/

func main() {
	result := helper.SayHello("Sintya")
	fmt.Println(result)
}