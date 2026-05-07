//go:build ignore

package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/**Package container/ring
adalah implementasi struktur data circular list
Circular list adalah struktur data ring, dimana diakhir element akan kembali
ke element awal (HEAD)

https://pkg.go.dev/container/ring
**/

func main() {
	// Buat ring
	var myRing = ring.New(5)

	// Isi data ke dalam ring
	for i := 0; i < myRing.Len(); i++ {
		myRing.Value = "Value:" + strconv.Itoa(i)
		myRing = myRing.Next() // Pindah ke elemen berikutnya
	}

	// Baca isi nilai di dlm ring dg fungsi Do(func(paramenter any/interface{}))
	// Ini akan berputar sekali melewati semua elemen
	myRing.Do(func(value any) {
		fmt.Println(value)
	})

	// Membuktikan ini sirkular:
	// Jika kita Next() 3 kali dari elemen 1, kita kembali ke elemen 1
	node1 := myRing
	node5 := myRing.Next().Next().Next().Next().Next()

	fmt.Println(node1 == node5) // Hasilnya: true

}
