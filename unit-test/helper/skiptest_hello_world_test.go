package helper

import (
	"runtime"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	// GOOS (GO Operating System)
	if runtime.GOOS == "linux" {
		// membatalkan test
		t.Skip("Can't run on Linux")
	}

	result := HelloWorld("Sintya")
	// Pengecekan
	require.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")
}
