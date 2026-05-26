package entity

/*
  - Mock adalah Membuat object palsu agar unit test fokus ke logic yang sedang dites.
  - Mock merupakan salah satu teknik dlm unit testing, dimana kita bisa
    membuat mock object dari suatu object yg memang sulit untuk ditesting
  - testify mendukung pembuatan mock object ini, namun perlu diperhatikan
    jika desain program kita jelek, akan sulit untuk melakukan mocking.

Contoh Kasus:
* Membuat contoh aplikasi golang yang melakukan query ke database.
* Dimana kita akan buat layer Service sbg business logic, &
  layer Repository sbg jembatan ke database.
* Agar kode kita mudah dites, disarankan agar membuat kontrak
  berupa interface.

https://github.com/stretchr/testify
https://pkg.go.dev/github.com/stretchr/testify/mock
*/

// Struct ini dibuat untuk representasi isi dari table database yg digunakan
type Category struct {
	Id   string
	Name string
}
