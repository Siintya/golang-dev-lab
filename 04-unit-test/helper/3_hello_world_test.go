package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldSintya(t *testing.T) {
	result := HelloWorld("Sintya")

	if result != "Hello Sintya" {
		// panic("Result is not 'Hello Sintya'")

		// error
		t.Error("Result must be 'Hello Sintya'")
	}

	fmt.Println("TestHelloWorld is DONE")
}

func TestHelloWorldAjeng(t *testing.T) {
	result := HelloWorld("Lestari")

	if result != "Hello Lestari" {
		// panic("Result is not 'Hello Lestari'")

		// error
		t.Fatal("Result must be 'Hello Ajeng'")
	}

	fmt.Println("TestHelloWorldLestari is DONE")
}
