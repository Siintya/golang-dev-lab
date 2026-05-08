//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**File Management
https://chmod-calculator.com/
**/

// READ FILE sample.go
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	result, _ := readFile("sample.log")
	fmt.Println(result)
}
