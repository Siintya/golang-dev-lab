//go:build ignore

package main

import (
	"bufio"
	"os"
)

/**Package io (Input/Output)
* merupakan fitur di golang yg digunakan sbg standard untuk
  proses input output.
* di golang, semua mekanisme input output pasti mengikuti strandard package io
https://pkg.go.dev/io
**/

// untuk membaca input, golang menggunakan kontrak interface yg bernama

func main() {
	// Writer
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world \n")
	_, _ = writer.WriteString("Selamat Belajar")
	writer.Flush()
}
