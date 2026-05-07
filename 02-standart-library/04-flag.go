//go:build ignore

package main

import (
	"flag"
	"fmt"
)

/**Package flag
berisi fungsionalitas untuk memparsing command line argument
https://pkg.go.dev/flag
**/

func main() {
	var username = flag.String("username", "root", "database username")
	var password = flag.String("password", "root", "database password")
	var host = flag.String("host", "localhost", "database host")
	var port = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	/**Jika ingin mengganti  masukan parameter berikut ini, ketika
	run file golang:
	$ go run 04-flag.go -username=sintya -password=rahasia -host=123.123.123.3 -port=5505
	**/

}
