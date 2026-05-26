package helper

import (
	"testing"
)

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Sintya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sintya")
		}
	})

	b.Run("Lestari", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lestari")
		}
	})
}

/*
Testing hanya Benchmark tanpa unit test
$ go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorldSub

>> Output:
Sebelum Unit Test
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkHelloWorldSub
BenchmarkHelloWorldSub/Sintya
BenchmarkHelloWorldSub/Sintya-4                 49693666                23.82 ns/op
BenchmarkHelloWorldSub/Lestari
BenchmarkHelloWorldSub/Lestari-4                50109606                23.74 ns/op
PASS
Setelah Unit Test
ok      unit-test/helper        2.439s
===========================================

$ go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorldSub/Sintya

>> Output:
Sebelum Unit Test
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkHelloWorldSub
BenchmarkHelloWorldSub/Sintya
BenchmarkHelloWorldSub/Sintya-4                 43564638                24.12 ns/op
PASS
Setelah Unit Test
ok      unit-test/helper        1.100s

*/
