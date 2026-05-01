package main

import "fmt"

/**Nil (Data Kosong)
* nil adalah kata kunci yg mewakili nilai kosong atau "ketiadaan nilai" (zero value) untuk tipe-tipe data tertentu yang bersifat referensi atau pointer.
* Tidak semua tipe data di Go bisa diisi dengan nil. 
  berikut ini beberapa tipe data yg dapat digunakan nil.
  >> pointer: Penanda bahwa pointer tidak menunjuk ke alamat memori mana pun.
  >> slice: Menunjukkan slice yang belum diinisialisasi (tidak memiliki array pendukung).
  >> map: Map yg dideklarasikan tapi belum dibuat dg make.
  >> channel: Jalur komunikasi yang belum dibuka.
  >> interface: Interface yang tidak membawa nilai maupun tipe konkret.
  >> function: Variabel tipe fungsi yang belum diisi dengan fungsi tertentu.
**/

// Contoh nil pada map
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	// Eksekusi
	data := NewMap("Sintya")
	
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}