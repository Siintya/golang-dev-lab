//go:build ignore

package main

import "os"

/**File Management
https://chmod-calculator.com/
**/

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.log", "this is sample log")
}
