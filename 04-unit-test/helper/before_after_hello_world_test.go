package helper

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum Unit Test")

	// eksekusi semua unit test
	m.Run()

	// after
	fmt.Println("Setelah Unit Test")
}
