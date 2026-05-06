package main

import (
	"math"
	"fmt"
)

/**Package math
berisi constant dan fungsi matematika
https://pkg.go.dev/math
**/

func main() {
	// math.Round(float64): membulatkan nilai float64 keatas atau kebawah, sesuai dg yg paling dekat
	fmt.Println("Round:", math.Round(1.69)) // 2
	
	// math.Floor(float64): membulatkan float64 kebawah
	fmt.Println("Floor:", math.Floor(1.99)) // 1
	
	// math.Ceil(float64): membulatkan float64 keatas
	fmt.Println("Ceil:", math.Ceil(1.15)) // 2
	
	// math.Max(float64): mengembalikan nilai float64 paling besar
	fmt.Println("Max:", math.Max(1, 20)) // 20
	
	// math.Min(float64): mengembalikan nilai float64 paling kecil
	fmt.Println("Min:", math.Min(1, 20)) // 1
}



