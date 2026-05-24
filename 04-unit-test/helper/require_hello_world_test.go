package helper

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sintya")

	// Jika pengecekan gagal, akan memanggil FailNow()
	require.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")

	fmt.Println("TestHelloWorld with Require is DONE")
}
