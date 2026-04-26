package main

import "fmt"

func main() {
	var grade = 95
	var attandance = 80

	var finalPassedGrade bool = grade > 80
	var passedAttandance bool = attandance > 80

	// && akan bernilai true jika semua nilainya TRUE
	var passed bool = finalPassedGrade && passedAttandance
	fmt.Println(passed)

	// || (OR) akan bernilai false jika semua nilainya FALSE
	var passedRemidial bool = finalPassedGrade || passedAttandance
	fmt.Println(passedRemidial)

	// ! (NOT)
	var passed2 bool = !passedRemidial
	fmt.Println(passed2)
}