//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**Package io (Input/Output)
* merupakan fitur di golang yg digunakan sbg standard untuk
  proses input output.
* di golang, semua mekanisme input output pasti mengikuti strandard package io
https://pkg.go.dev/io
**/

// untuk membaca input, golang menggunakan kontrak interface yg bernama

func main() {
	// Reader
	input := strings.NewReader("this is long string \nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}
