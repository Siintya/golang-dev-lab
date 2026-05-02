package main

import (
	"01-fundamentals/database"
	"fmt"
)

/**Package Initialization
* package ini merupakan proses "persiapan" atau "pemanasan" yg dilakukan
  Golang sebelum program utama kita benar-benar berjalan.
* init() adalah fungsi spesial yang akan dijalankan secara otomatis. 
  kita tdk perlu memanggilnya secara manual; Golang yg akan 
  mengeksekusinya untuk kita.
* Aturan utama fungsi init():
  >> berjalan scr otomatis, dieksekusi sebelum fungsi main()
  >> dlm 1 file atau package, boleh memiliki >1 fungsi init()
  >> init() dijalankan sesuai urutan deklarasinya.
  >> init() hanya dijalankan 1 kali selama program hidup, tepat saat
     package tsb pertama kali di-load.
* Manfaat proses initialization ini:
  >> mengecek apakah koneksi ke database sudah siap.
  >> membaca file konfigurasi (seperti username/password) 
     sebelum program mulai.
  >> mengisi variabel yang nilainya tidak bisa ditulis langsung scr 
     sederhana.
**/

func main() {
	fmt.Println(database.GetDatabase())
}