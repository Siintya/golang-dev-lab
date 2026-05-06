package main

import (
	"time"
	"fmt"
)

/**Package time
berisi fungsionalitas untuk management waktu di golang
https://pkg.go.dev/time
**/

func main() {
	// time.Now(): untuk mendapatkan waktu saat ini
	now := time.Now()
	fmt.Println("Waktu saat ini:", now.Local())

	// time.Date(...): untuk membuat waktu
	utc := time.Date(2026, time.May, 4, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// time.Parse(layout, string): untuk memparsing waktu dari string
	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(parse)

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Day(), valueTime.Month(), valueTime.Year())
}