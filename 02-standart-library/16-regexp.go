//go:build ignore

package main

import (
	"fmt"
	"regexp"
)

/*
*Package Regexp (Regular Expression)
* adalah sebuah pola (pattern) yg digunakan
  untuk mencari, mencocokkan, dan memanipulasi teks.
* Regexp memungkinkan kita untuk mencari string yg rumit
  tanpa harus menulis banyak logika if-else manual.
* Go menyediakan paket bawaan bernama regexp yg
    sangat efisien dan aman (menggunakan sintaks RE2).
* Kapan Menggunakan Regexp?
    >> Validasi Format Data
    >> Pencarian Teks Kompleks (Pattern Matching)
    >> Ekstraksi Data (Data Scraping)
    >> Pembersihan dan Penggantian Teks (Search & Replace)

https://github.com/google/re2/wiki/Syntax
https://golang.org/pkg/regexp
*/

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])a`)

	fmt.Println(regex.MatchString("ana"))
	fmt.Println(regex.MatchString("aca"))
	fmt.Println(regex.MatchString("Ara"))

	fmt.Print(regex.FindAllString("aca ana ama ani a2a ara", 2))
}
