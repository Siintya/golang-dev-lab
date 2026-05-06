package main

import (
	"strconv"
	"fmt"
)

/**Package strconv (String Conversions)
https://pkg.go.dev/strconv
**/

func main() {
	// strconv.parseBool(string): mengubah string ke bool
	result, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println("Strconv.ParseBool:", result)
	} else {
		fmt.Println("Error", err.Error())
	}

	// strconv.parseFloat(string, bitSize int): mengubah string ke float
	result2 := "3.14"

	if res, err := strconv.ParseFloat(result2, 32); err == nil {
		fmt.Printf("Strconv.ParseFloat: %v\n", res)
	} else {
		fmt.Println("Error", err.Error())
	}


	// strconv.parseInt(string, base int, bitSize int): mengubah string ke int64
	result3, err := strconv.ParseInt("14500", 16, 64)
	if err == nil {
		fmt.Println("Strconv.ParseInt:", result3)
	} else {
		fmt.Println("Error", err.Error())
	}

	// strconv.FormatBool(string): mengubah bool ke string
	boolean := strconv.FormatBool(true)
	fmt.Println("strconv.FormatBool:", boolean)

	// strconv.FormatFloat(float,...): mengubah float64 ke string
	result4 := strconv.FormatFloat(4.587839784, 'e', -1, 32)
	fmt.Println("strconv.FormatFloat:", result4)

	// strconv.FormatInt(int,...): mengubah int64 ke string
	result5 := strconv.FormatInt(8899, 10)
	fmt.Println("strconv.FormatInt:", result5)

	// Numeric Conversions
	result6, err := strconv.Atoi("15")
	if err == nil {
		fmt.Println("Atoi:", result6)
	}

	result7 := strconv.Itoa(90)
	fmt.Printf("Itoa: %T, %v\n", result7, result7)
}



