//go:build ignore

package main

import (
	"encoding/csv"
	"os"
)

/*
*Package Encoding dg CSV Writer
* CSV (Comma Separated Value) jenis extention dokumen yg biasa digunakan
  di excel.
https://pkg.go.dev/encoding/json
*/

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"ajeng", "sintya", "lestari"})
	_ = writer.Write([]string{"budi", "pratama", "saputra"})
	_ = writer.Write([]string{"joko", "deva", "anwar"})
	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})

	writer.Flush()
	/*Output:
	ajeng,sintya,lestari
	budi,pratama,saputra
	joko,deva,anwar
	eko,kurniawan,khannedy
	*/
}
