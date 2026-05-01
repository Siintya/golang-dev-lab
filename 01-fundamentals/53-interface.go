package main

import "fmt"

/**Interface
* Adalah tipe data yg abstract, dia tdk memiliki implementasi data scr langsung
* Sebuh interface berisikan defini-definisi method
* Biasanya interface digunakan sbg kontrak, berarti harus ada yg meng-implementasikan.
* Biasanya interface itu diimplemetasikan dlm bentuk sebuah struct
* Implementasi interface
  >> setiap tipe data yg sesuai dg kontrak interface, scr otomatis dianggap sbg interface tsb
  >> Sehingga kita tidak perlu mengimplementasikan interface scr manual
  >> Hal ini sedikit berbeda dg bahasa pemrograman lain (Java) yg ketika membuat interface,
     kita harus menyebutkan scr eksplisit untuk menggunakan interface yg akan diimplementasikan.
**/

type HasName interface {
	// method GetName()
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

// Implementasi dari Interface
func (person Person) GetName() string {
	return person.Name
}

// Implementasi dari Interface
type Pet struct {
	Name string
}

func (pet Pet) GetName() string {
	return pet.Name
}

func main() {
	person := Person{Name: "Sintya"}
	SayHello(person)

	pet := Pet{Name:"Kucing"}
	SayHello(pet)

}