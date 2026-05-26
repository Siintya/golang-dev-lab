package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Table Test: dimana kita menyediakan data berupa slice yang berisi parameter dan
ekspektasi hasil dari unit test.
*/
func TestTableTest(t *testing.T) {
	// 1. Buat table array test
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Sintya",
			request:  "Sintya",
			expected: "Hello Sintya",
		},
		{
			name:     "Lestari",
			request:  "Lestari",
			expected: "Hello Lestari",
		},
	}

	// 2. Looping
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
