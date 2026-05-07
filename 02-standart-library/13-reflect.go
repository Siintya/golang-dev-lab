package main

import (
	"reflect"
)

/**Package reflect (reflection)
* dlm bahasa pemrograman, biasanya ada fitur Reflection, dimana
  kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
* hal ini bisa dilakukan di golang dg menggunakan package reflect
* Reflection sangat berguna ketika kita ingin membuat library
  yg general sehingga mudah digunakan
  https://pkg.go.dev/reflect
**/

type Sample struct {
	Name string
}

type Person struct {
	Name string
}

func readFile(value array) {
	var valueType reflect.Value = reflect.TypeOf(value)
}

func main() {

}
