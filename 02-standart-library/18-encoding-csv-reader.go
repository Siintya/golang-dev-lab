//go:build ignore

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

/*
*Package Encoding dg CSV Reader
* CSV (Comma Separated Value) jenis extention dokumen yg biasa digunakan
  di excel.
https://pkg.go.dev/encoding/json
*/

func main() {
	csvString := "ajeng,sintya,lestari\n" +
		"budi,pratama,saputra\n" +
		"joko,deva,anwar\n" +
		"eko,kurniawan,khannedy"

	// encode: mengubah data program menjadi format tertentu
	reader := csv.NewReader(strings.NewReader(csvString))

	// jika perulangan seperti ini artinya perungan tidak akan berhenti
	// karna tdk ada condition disini
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
		/*Output:
		[ajeng sintya lestari]
		[budi pratama saputra]
		[joko deva anwar]
		[eko kurniawan khannedy]
		*/
	}

	// decode: mengubah format tersebut kembali menjadi data program
	// Contoh decode dg base64

}
