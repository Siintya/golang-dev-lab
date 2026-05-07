//go:build ignore

package main

import (
	"fmt"
	"reflect"
)

/*
*StructTag - Validation Input
Disini akan mengimplementasikan strctTag pada Validation input
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

func IsValid(value interface{}) (result bool) {
	result = true
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// validasi
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return true
}

func main() {
	// running
	person := Person{
		Name:    "Sintya",
		Address: "Bekasi",
		Email:   "ada",
	}

	fmt.Println(IsValid(person)) // output: true
}
