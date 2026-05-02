package main

import (
	"fmt"
	"os"
)

/**Package OS-https://pkg.go.dev/os@go1.26.2
* Berisi fungsionalitas untuk mengakses fitur sistem operasi scr 
  independen (bisa digunakan disemua OS)
* Args -> arguments
**/

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

}