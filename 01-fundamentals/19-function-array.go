package main

import "fmt"

func main() {
	var values = [...]int{
		80,
		70,
		65,
		100,
	}
	fmt.Println(values)

	// len(array): untuk mendapatkan panjang array
	fmt.Println(len(values))

	// array[index]: mendapatkan data di posisi index
	fmt.Println(values[2])

	// array[index] = value: mengubah data diposisi index
	values[2] = 20
	fmt.Println(values[2])
}