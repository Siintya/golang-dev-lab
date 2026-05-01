package main

import "fmt"

/**Struct Method
* Struct sebagai parameter fungsi
  struct tidak ada bedanya dengan int atau string dalam hal pengiriman data.
  Kita bisa memasukkan struct ke dalam fungsi biasa.
* Menambahkan method ke struct
  kita bisa "menempelkan" fungsi tersebut langsung ke struct-nya. 
* Method adalah function yg menempel di dlm struct
**/
type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "! My name is", customer.Name)
}

func main() {
	sintya 	:= Customer{"Sintya", "Indonesia", 25}
	dewi 	:= Customer{"Dewi", "Indonesia", 25}

	// Fungsi sayHello ini bisa diakses ketika sudah membuat object struct
	sintya.sayHello("Lestari")
	dewi.sayHello("Lestari")
}