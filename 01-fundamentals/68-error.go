package main

import (
  "errors"
  "fmt"
)

/**Error Interface
* golang memiliki interface yg digunakan sbg kontrak untuk membuat error,
  nama interface-nya adalah error.
* the error built-in interface type is the conventional interface for
* representing an error condition, with the nil value representing 
  no error.
**/
// Membuat Error
// golang sudah menyediakan library untuk membuat helper scr mudah, yg
// terdapat di package errors

func Pembagian(nilai int, pembagi int) (int, error) {
  if pembagi == 0 {
    // membuat error errors.New("...")
    return 0, errors.New("Pembagian dengan 0")
  } else {
    return nilai / pembagi, nil
  }
}

func main() {
  hasil, err := Pembagian(100, 0)
  if err == nil {
    fmt.Println("Hasil", hasil)
  } else {
    fmt.Println("Error", err.Error())
  }
}