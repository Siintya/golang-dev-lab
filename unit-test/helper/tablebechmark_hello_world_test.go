package helper

import (
	"testing"
)

/*
Programmer Golang terbiasa membuat table benchmark agar mudah dalam
melakukan performance test dengan kombinasi data yg berbeda-beda tanpa harus
membuat banyak benchmark function
*/
func BenchmarkTable(b *testing.B) {
	// 1. Buat table array test
	benchmarks := []struct {
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
		{
			name:     "Ajeng Shintya Lestari",
			request:  "Ajeng Shintya Lestari",
			expected: "Hello Ajeng Shintya Lestari",
		},
	}

	// 2. Looping
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

/*
$ go test -v -run=TestTidakAda -bench=BenchmarkTable

>> Output:
Sebelum Unit Test
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkTable
BenchmarkTable/Sintya
BenchmarkTable/Sintya-4                 48649764                23.89 ns/op
BenchmarkTable/Lestari
BenchmarkTable/Lestari-4                49897962                24.07 ns/op
BenchmarkTable/Ajeng_Shintya_Lestari
BenchmarkTable/Ajeng_Shintya_Lestari-4          44343704                24.54 ns/op
PASS
Setelah Unit Test
ok      unit-test/helper        3.538s
*/
