package main

import "fmt"

func main() {
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daySlice1 := days[5:]
	// fmt.Println(daySlice1)

	// append(slice, data) -> membuat slice baru dg menambah data ke posisi terakhir slice, jika kapasitas sudah penuh maka akan membuat array baru.
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	// fmt.Println(daySlice1)
	// fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	// fmt.Println(daySlice1)
	// fmt.Println(daySlice2)
	// fmt.Println(days)

	// make([]TypeData, length, capacity) -> membuat slice baru
	newSlice:= make([]string, 2, 5)
	newSlice[0] = "Sintya"
	newSlice[1] = "ajeng"
	// fmt.Println(newSlice)

	// len(slice) -> untuk mendapatkan panjang
	// fmt.Println(len(newSlice))

	// cap(slice) -> untuk mendapatkan kapasitas
	// fmt.Println(cap(newSlice))

	// newSlice2 := append(newSlice, "Sintya")
	// fmt.Println(newSlice2)
	// fmt.Println(len(newSlice2))
	// fmt.Println(cap(newSlice2))

	// copy(destination, source) -> menyalin slice dari souce ke destination
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}