package main

import _ "01-fundamentals/internal"

/**Blank Indetifier (_)
* Blank Identifier adalah sebuah simbol garis bawah (_) yg 
  digunakan untuk memberi tahu compiler Go: "Saya tahu ada 
  data atau sesuatu di sini, tapi saya tidak butuh dan 
  tolong abaikan saja."
* Di bahasa pemrograman lain, kamu mungkin bisa membuat variabel tapi 
  tidak pernah memakainya. Namun, di Go, itu dilarang keras. 
  Jika kita membuat variabel atau mengimpor package tapi tidak 
  menggunakannya, program kita tidak akan bisa jalan (error).
  Karena aturan Go yang sangat "disiplin" ini, kita butuh _ sebagai 
  tempat pembuangan untuk hal-hal yang tidak kita perlukan.
* Menampung nilai yang wajib diterima tapi tidak akan kita gunakan.
**/

func main() {

}