//go:build ignore

package main

import (
	"encoding/base64"
	"fmt"
)

/*
*Package Encoding
* sebuah standar interface dan kumpulan sub-paket yg digunakan untuk
  mengubah data dari satu bentuk ke bentuk lain.
  Proses ini biasanya disebut sebagai Encoding dan Decoding.
* yg paling populer: base64, csv & json
* Kita akan butuh package ini setiap kali data harus "keluar"
  atau "masuk" ke program kita:
  >> Interaksi API
  >> Penyimpanan File
  >> Keamanan/Web
  >> Database
https://pkg.go.dev/encoding/json
*/

func main() {
	// encode: mengubah data program menjadi format tertentu
	// Contoh encode dg base64
	value := "Sintya Lestari"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// decode: mengubah format tersebut kembali menjadi data program
	// Contoh decode dg base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
