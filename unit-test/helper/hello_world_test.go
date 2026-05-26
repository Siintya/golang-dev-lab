package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sintya")

	if result != "Hello Sintya" {
		// panic("Result is not 'Hello Sintya'")

		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorld is DONE")
}

func TestHelloWorldLestari(t *testing.T) {
	result := HelloWorld("Lestari")

	if result != "Hello Lestari" {
		// panic("Result is not 'Hello Lestari'")

		// error
		t.FailNow()
	}

	fmt.Println("TestHelloWorldLestari is DONE")
}
