//go:build ignore

package main

import (
	"fmt"
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
	Name, Address, Email string
	Age                  int
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func main() {
	readField(Sample{"Sintya"}) // Type Name Sample, Name with type string
	/**Output:
	Type Name Sample
	Name with type string
	**/

	readField(Person{"Sintya", "", "", 20})
	/**Output:
	Type Name Person
	Name with type string
	Address with type string
	Email with type string
	Age with type int
	**/
}
