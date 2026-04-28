package main

import "fmt"

// **Named Return Values
//  Mendeklarasikan tipe data return value di function
func getData() (firstName string, middleName string, lastName string, age int) {
	firstName 	= "Ajeng"
	middleName 	= "Sintya"
	lastName 	= "Lestari"
	age = 25

	return firstName, middleName, lastName, age
}

func main() {
	firstName, middleName, lastName, age := getData()
	fmt.Println("Name:", firstName, middleName, lastName, "\nAge:", age)

}