package main

import "fmt"

/**Interface Kosong
* interface{} adalah interface yg tidak memiliki dklarasi method satupun, hal ini membuat
  scr otomatis semua tipe data akan menjadi implementasi-nya
* Karena interface kosong tidak meminta method apa pun, maka secara otomatis semua tipe data (int, string, struct, boolean, dll.) 
  memenuhi syarat sebagai interface kosong.
* interface kosong, jga memiliki type alias bernama any
**/
func Ups() any {
	// return 1
	// return true
	// return "Hello"
	return 0.5
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}