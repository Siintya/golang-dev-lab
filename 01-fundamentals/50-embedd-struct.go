package main

import "fmt"

type ContactInfo struct {
	Phone   string
	Address string
}

type Admin struct {
	Username    string
	Level       int
	ContactInfo // Ini adalah Embedding (Anonymous Field)
}

func main() {
	adm := Admin{
		Username: "admin01",
		Level:    1,
		ContactInfo: ContactInfo{
			Phone:   "08123456789",
			Address: "Bekasi",
		},
	}

	// Kita bisa langsung mengakses field dari ContactInfo
	// lewat objek Admin tanpa harus menulis adm.ContactInfo.Address
	fmt.Println(adm.Username) // Output: admin01
	fmt.Println(adm.Address)  // Output: Bekasi (Direct Access karena embedding)

	// Akses secara lengkap juga tetap bisa dilakukan:
	fmt.Println(adm.ContactInfo.Phone) // Output: 08123456789
}
