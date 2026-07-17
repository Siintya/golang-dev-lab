package main

import "fmt"

func main() {
	message := ` This is a raw string literal.
	It can span multiple lines
	without escape sequences like \n.`
	fmt.Println(message)
}
