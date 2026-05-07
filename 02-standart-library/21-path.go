//go:build ignore

package main

import (
	"fmt"
	"path"
)

/**Package Path
* digunakan untuk memanipulasi data path seperti path di URL
  atau path di File
* scr default, package ini menggunakan slash (/) sbg karakter path-nya.
  Oleh karna itu, cocok untuk data URL
* Namun, jika ingin menggunakan untuk memanipulasi di File System, karna
  OS Windows menggunakan backslash (\), maka khusus untuk File System tsb
  perlu menggunakan package path/filepath
https://pkg.go.dev/path
https://pkg.go.dev/path/filepath
**/

func main() {
	// Directory atau Folder
	fmt.Println(path.Dir("hello/world.go")) //hello
	// File
	fmt.Println(path.Base("hello/world.go")) // world.go
	// Ext -> Extention
	fmt.Println(path.Ext("hello/world.go"))             // .go
	fmt.Println(path.Join("hello", "world", "main.go")) //hello/world/main.go
}
