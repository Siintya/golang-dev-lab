//go:build ignore

package main

import (
	"fmt"
	"sort"
)

/**Package sort
- berisi utilitas untuk proses pengurutan data.
- sebelum mengurutkan data, yg dilakukan adalah mengimplementasikan
  kontrak di interface sort.interface
https://pkg.go.dev/sort
**/
// dua cara menggunakan package ini
// 1. Mengurutkan Tipe Data Dasar
func SortData(data any) {
	switch value := data.(type) {
	case []int:
		sort.Ints(value)
		fmt.Println("Sorted Ints:", value)
	case []string:
		sort.Strings(value)
		fmt.Println("Sorted String:", value)
	default:
		fmt.Println("Type data not supported for sorting!")
	}
}

// 2. mengimplementasikan interface untuk tipe data yg lebih kompleks.
type User struct {
	Name string
	Age  int
}

// menempelkan (attach) method yang dibutuhkan oleh interface sort.
type UserSlice []User

// Ini memberi tahu package 'sort' seberapa banyak data yg perlu diperiksa.
// Len mengembalikan jumlah elemen dalam slice.
func (s UserSlice) Len() int {
	return len(s)
}

// Di sini, mengurutkan berdasarkan Age secara Ascending.
// Less menentukan aturan pengurutan (logika perbandingan).
// Mengembalikan true jika elemen pada index i harus berada di depan index j.
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// Method ini dipanggil oleh algoritma sortir ketika Less(i, j) bernilai false
// dan posisi elemen perlu dipindahkan.
// Swap bertugas untuk menukar posisi dua elemen dalam slice.
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	angka := []int{5, 2, 9, 1, 5, 6}
	SortData(angka) // Output: Sorted Ints: [1 2 5 5 6 9]

	buah := []string{"Jeruk", "Mangga", "Rambutan"}
	SortData(buah) // Output: Sorted String: [Jeruk Mangga Rambutan]

	users := []User{
		{"Rudi", 27},
		{"Tono", 24},
		{"Nana", 26},
		{"Dewi", 20},
	}

	// melakukan konversi tipe dari []User ke UserSlice agar
	// interface sort.Sort bisa mengenali method Len, Less, dan Swap.
	sort.Sort(UserSlice(users))
	fmt.Println(users) // Output: [{Dewi 20} {Tono 24} {Nana 26} {Rudi 27}]

}
