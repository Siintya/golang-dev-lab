package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	Sub Test

Golang mendukung fitur pembuatan function unit test didalam
function unit test dinamakan sub test.
*/
func TestSubTest(t *testing.T) {
	t.Run("Sintya", func(t *testing.T) {
		result := HelloWorld("Sintya")
		require.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")
	})

	// Jika hanya ingin running test ini: $ go test -v -run=TestSubTest/Lestari
	t.Run("Lestari", func(t *testing.T) {
		result := HelloWorld("Lestari")
		require.Equal(t, "Hello Lestari", result, "Result must be 'Hello Lestari'")
	})
}
