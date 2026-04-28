package main

import "fmt"

// **Variadic Function
// Parameter yg ada di posisi terakhir, bisa dijadikan sebuah varargs
// varargs (variable arguments) yaitu datanya bisa menereima lebih dari satu input seperti array.
// array & varargs itu BERBEDA!
// Bedanya, jika parameter array harus buat array terlebih dahulu sebelum mengirimkan ke function.
// Sedangkan varargs, bisa langsung mengirim datanya, jika lebih dari satu cukup pake tanda koma.
// varargs -> ...datatype (contoh: ...int)

func sumAll(numbers ...int) int {
	total := 1
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println(total)
}