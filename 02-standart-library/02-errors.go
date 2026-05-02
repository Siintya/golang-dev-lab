package main

import (
	"fmt"
	"errors"
)

/**Package errors-https://pkg.go.dev/errors@go1.26.2**/

// #1 Create Error
var (
	ValidationError = errors.New("validation error")
	NotFoundError 	= errors.New("not found error")
)

// #2 Use Error
func getByUsername(username string) error {
	if username == "" {
		return ValidationError
	}

	if username != "sintya88" {
		return NotFoundError
	}

	// success!
	return nil
}

func main() {
	err := getByUsername("sintya88")

	// errors.Is() -> mengecek jenis type errornya
	if err != nil {	
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation ERROR!")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found ERROR!")
		} else {
			fmt.Println("Unknown ERROR!")
		}
	} else {
		fmt.Println("Success!")
	}
}