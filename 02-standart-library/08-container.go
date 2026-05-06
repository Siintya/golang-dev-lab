package main

import (
	"container/list"
	"fmt"
)

/**Package container/list
adalah implementasi struktur data Doubly linked list di golang

Apa itu Doubly Linked List?
* Berbeda dengan Array atau Slice yang menyimpan data dalam blok 
  memori yg berurutan, Doubly Linked List menyimpan data dlm unit-unit 
  terpisah yang disebut Node.
* Setiap Node dalam Doubly Linked List memiliki tiga komponen utama:
	- Value: Data yang disimpan (bisa berupa angka, string, atau objek).
	- Next: Pointer yang menunjuk ke Node selanjutnya.
	- Prev: Pointer yang menunjuk ke Node sebelumnya.
https://pkg.go.dev/container/list
**/

func main() {
	// Cara menggunakan Doubly Linked List
	var data = list.New()

	data.PushBack("Ajeng")
	data.PushBack("Sintya")
	data.PushBack("Lestari")

	var head *list.Element = data.Front();
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)

	next2 := next.Next()
	fmt.Println(next2.Value)

	for x := data.Front(); x != nil; x = x.Next() {
		fmt.Println("Use Looping:", x.Value)
	}

}



