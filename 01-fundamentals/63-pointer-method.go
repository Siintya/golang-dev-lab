package main

import "fmt"

/** Pointer di Method
* Walaupun method akan menempel di struct, tapi sebenarnya data stuct yg
  diakses didlm method tsb adalah pass by value.
* Jadi sangat disarankan menggunakan pinter di method untuk menghemat 
  memory karna harus slalu di-duplikasikan ketika memanggil method
**/
type Person struct {
	Name 	string
	Age 	int
}

func (person *Person) Married() {
	person.Name = "Mr. " + person.Name + ", Age"
	person.Age  = person.Age
}  

func main() {
	putra := &Person{"Putra", 28}
	putra.Married()

	fmt.Println(putra)
}