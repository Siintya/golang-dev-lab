package main

import "fmt"

// blok konstanta
const (
	StatusPending    = iota // otomatis bernilai 0
	StatusProcessing        // otomatis bernilai 1
	StatusCompleted         // otomatis bernilai 2
)

const (
	Monday  = iota // karna ganti blok const baru, iota RESET lagi dari 0
	Tuesday        // Otomatis bernilai 1
)

func main() {
	fmt.Println(StatusPending, StatusProcessing, StatusCompleted) // Output: 0 1 2
	fmt.Println(Monday, Tuesday)                                  // Output: 0 1
}
