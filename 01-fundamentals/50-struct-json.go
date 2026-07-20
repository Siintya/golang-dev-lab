package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	// Mengubah "FullName" menjadi "full_name" saat jadi JSON
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	// omitempty berarti field ini akan diabaikan jika nilainya kosong
	Password string `json:"password,omitempty"`
}

func main() {
	emp := Employee{
		FullName: "Shintya",
		Email:    "shintya@example.com",
	}

	// Mengubah struct menjadi JSON
	jsonData, _ := json.Marshal(emp)

	fmt.Println(string(jsonData))
	// Output: {"full_name":"Shintya","email":"shintya@example.com"}
	// Perhatikan: password tidak muncul karena kosong dan diberi tag omitempty
}
