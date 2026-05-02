package main

import (
	"01-fundamentals/helper"
	"fmt"
)

/**Access Modifier
* Yaitu sebuah aturan keamanan atau izin akses. Aturan ini menentukan 
  siapa sa yg boleh melihat atau menggunakan "sesuatu" (variabel, 
  fungsi, atau data) di dalam kode program kita.
* di golang, untuk menentukan access modifier cukup dg nama function
  atau variable.
* jika namanya (function) diawali dg huruf besar, maka artinya bisa
  diakses dari package lain, begitupun sebaliknya.
**/

func main() {
	sayhello := helper.SayHello("Sintya")
	fmt.Println(sayhello)

	fmt.Println(helper.Application)
	fmt.Println(helper.version)  // tdk bisa diakses dari package yg berbeda
	saygoodbye := helper.sayGoodBye("Tom")
	fmt.Println(saygoodbye) // tdk bisa diakses dari package yg berbeda
}