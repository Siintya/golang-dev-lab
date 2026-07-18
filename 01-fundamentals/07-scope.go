package main

import "fmt"

func main() {
	guest := "Sintya"

	if true {
		guest := "Budi"
		fmt.Println("Didalam IF", guest) // Output: Budi
	}

	fmt.Println("Diluar IF", guest) // Output: Sintya
}
