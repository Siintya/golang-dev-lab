//go:build ignore

package main

import (
	"os"
)

/**File Management
https://chmod-calculator.com/
**/

// READ FILE sample.go
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil

}

func main() {

	addToFile("sample.log", "\nThis is add message")
}
