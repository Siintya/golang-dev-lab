package helper

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sintya")

	/* Jika pengecekan gagal, akan memanggil FailNow()
	require.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")
	penjelasan:
	- t: menandai test gagal
	- "Hello Sintya": Nilai yang diharapkan, bisa berupa string, int, bool, struct, dll.
	- result: Nilai hasil sebenarnya dari program.
	- "Result must be 'Hello Sintya'": msgAndArgs... (optional), pesan tambahan jika test gagal.
	*/
	require.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")

	fmt.Println("TestHelloWorld with Require is DONE")
}
