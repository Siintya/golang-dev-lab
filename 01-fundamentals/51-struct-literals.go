package main

import "fmt"

/**Struct Literals
* Adalah cara atau penulisan sintaks dlm membuat sebuah data dari sebuah struct dan langsung memberikan nilai pada field-fieldnya secara bersamaan.
**/
type Customer struct {
	// Denifisikan field disini...
	Name, Address string
	Age 		  int
}

func main() {
	// Cara 1
	budi := Customer{
		Name: 		"Budi",
		Address: 	"Indonesia",
		Age: 		28,
	}
	fmt.Println("Customer 1:", budi)

	// Cara 2
	dewi := Customer{"Dewi", "Indonesia", 30}
	fmt.Println("Custome 2:", dewi)
}