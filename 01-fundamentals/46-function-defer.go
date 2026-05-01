package main

import "fmt"

// **Defer Function
// # Adalah function yg bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai dieksekusi.
// # Defer function akan slalu dieksekusi meski terjadi error.
// # Defer menjaga logika "Buka" & "Tutup" tetap berdekatan secara visual di dalam kode. Ini membantu developer menghindari bug klasik di mana mereka lupa menutup sumber daya (file/koneksi) karena teralihkan oleh logika pemrosesan data yang panjang di tengah fungsi.
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // ditunda
	fmt.Println("Run Application") // dieksekusi langsung
}

func main() {
	runApplication()
}