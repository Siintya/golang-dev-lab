//go:build ignore

package main

import (
	"fmt"
	"reflect"
)

/*
*StructTag
  - Fitur structTug adalah string opsional yang diletakkan
    setelah tipe data dalam deklarasi field sebuah struct.
  - Tag ini berfungsi sebagai metadata yang memberikan
    instruksi tambahan kepada paket atau library
    lain tentang cara memproses field tersebut.
  - Cara penhulisanya: Tag ditulis di dalam
    tanda backtick (`) dan biasanya mengikuti format key:"value".
  - Kita akan sering menjumpai StructTag dalam skenario berikut:
    >> Serialisasi Data (JSON, XML, YAML)
    >> Pemetaan Database (ORM)
    >> Validasi Input
    >> Pemetaan Form Web

*
*/
type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)

		fmt.Printf("Field: %s\n", structField.Name)
		fmt.Printf("  - Required: %s\n", structField.Tag.Get("required"))
		fmt.Printf("  - Max:      %s\n", structField.Tag.Get("max"))
	}
}

func main() {
	// running
	readField(Sample{"Sintya"})
	readField(Person{"Sintya", "Bekasi", "ajeng.sintya15@gmail.com"})
}
