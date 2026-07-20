package main

import "fmt"

func main() {
	/**Slice
	 # adalah tipe data yang merepresentasikan potongan dari array.
	 # Berbeda dengan array, slice tidak memiliki panjang tetap sehingga ukurannya dapat berubah (dinamis).
	**/
	fruits := [...]string{
		"Banana",
		"Mango",
		"Strawberry",
		"Pineapple",
		"Grapes",
		"Orange",
		"Watermelon",
	}

	// Cara membuat slice dari array:
	// 1. array[low:high] -> membuat slice dari array dimulai index low sampai index sebelum high
	slice1 := fruits[3:7]
	fmt.Println(slice1)

	// 2. array[low:] -> membuat sclice dari array dimulai index low sapai index akhir di array
	slice2 := fruits[:3]
	fmt.Println(slice2)

	// 3. array[:high] -> membuat slice dari array dimulai index 0 sampai index akhir di array
	slice3 := fruits[3:]
	fmt.Println(slice3)

	// 4. array[:] -> membuat slice dari array dimulai index 0 sampai index akhir di array
	slice4 := fruits[:]
	fmt.Println(slice4)

	// make() -> membuat slice baru dengan panjang dan kapasitas tertentu
	score := []int{5, 10, 20}
	fmt.Println("Skor saat ini", score)

	score2 := make([]int, 3) // membuat slice baru dengan panjang 3 dan kapasitas 3
	fmt.Println("Skor saat ini", score2)

	// 1. Membuat slice lewat "Slice Literal"
	buah := []string{"Apel", "Jeruk", "Mangga"}
	fmt.Printf("Awal  -> Len: %d, Cap: %d, Data: %v\n", len(buah), cap(buah), buah)
	// Output: Len: 3, Cap: 3, Data: [Apel Jeruk Mangga]

	// 2. Terjadi "Growth" (Pertumbuhan) karena kapasitas awal (3) sudah penuh
	// Go secara otomatis mencarikan tempat memori baru yang kapasitasnya lebih besar
	buah = append(buah, "Pisang")
	fmt.Printf("Baru  -> Len: %d, Cap: %d, Data: %v\n", len(buah), cap(buah), buah)
	// Output: Len: 4, Cap: 6, Data: [Apel Jeruk Mangga Pisang]

	sliceData := []string{"Go", "Java", "Python"}

	// 1. Konversi langsung ke Array (Fitur Go 1.20+)
	// Ukuran array harus sama dengan atau lebih kecil dari panjang slice
	arr := [3]string(sliceData)
	fmt.Println(arr) // Output: [Go Java Python]

	// 2. Konversi ke Pointer Array
	arrPtr := (*[3]string)(sliceData)
	fmt.Println(*arrPtr) // Output: [Go Java Python]
}
