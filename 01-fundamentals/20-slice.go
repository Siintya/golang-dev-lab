package main

import "fmt"

func main() {
	// Slice adalah tipe data yang merepresentasikan potongan dari array. 
	// Berbeda dengan array, slice tidak memiliki panjang tetap sehingga ukurannya dapat berubah (dinamis).	
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
}