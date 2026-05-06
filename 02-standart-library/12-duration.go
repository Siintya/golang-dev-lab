package main

import (
	"time"
	"fmt"
)

/**Package time - Duration
* saat menggunakan tipe data waktu, kadang membutuhkan data berapa durasi
* package time memiliki type Duration, yg sebenarnya adalah alias untuk int64.
* namun terdapat banyak method yg bisa kita gunakan untuk memanipulasi duration.
https://pkg.go.dev/time
**/

func main() {
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration1)
	fmt.Println(duration2)
	fmt.Println(duration3)
	fmt.Printf("duration1 : %d\n", duration1)
	fmt.Printf("duration2 : %d\n", duration2)
	fmt.Printf("duration3 : %d\n", duration3)
}