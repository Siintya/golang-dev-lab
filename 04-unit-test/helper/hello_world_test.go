package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sintya")

	if result != "Hello Sintya" {
		panic("Result is not 'Hello Sintya'")
	}
}

func TestHelloWorldLestari(t *testing.T) {
	result := HelloWorld("Lestari")

	if result != "Hello Lestari" {
		panic("Result is not 'Hello Lestari'")
	}
}
