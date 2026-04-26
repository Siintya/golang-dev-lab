package main

import "fmt"

func main() {
	// **Function Map
	// make(map[TypeKey]TypeValue) -> membuat map baru
	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "Ups!"
	fmt.Println(book)
	
	// len(map) -> untuk mendapatkan jumlah data di map
	fmt.Println(len(book))

	// map[key] -> mengambil data di map dengan key
	fmt.Println(book["title"])

	// map[key] = value -> mengubah data di map dengan key
	book["author"] = "Sintya"
	fmt.Println(book)

	// delete(map, key) -> menghapus data di map dg key
	delete(book, "wrong")
	fmt.Println(book)


}