package helper

import (
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sintya")
	}
}

func BenchmarkHelloWorldLestari(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lestari")
	}
}

/*
Testing hanya Benchmark tanpa unit test
$ go test -v -run=NotMathUnitTest -bench=.

>> Output:
Sebelum Unit Test
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkHelloWorld
BenchmarkHelloWorld-4           46861101                23.99 ns/op
BenchmarkHelloWorldLestari
BenchmarkHelloWorldLestari-4    49857896                23.97 ns/op
PASS
Setelah Unit Test
===================

$ go test -v -run=NotMathUnitTest -bench=

>> Output:
Sebelum Unit Test
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkHelloWorldLestari
BenchmarkHelloWorldLestari-4    50402559                23.76 ns/op
PASS
Setelah Unit Test
ok      unit-test/helper        2.225s
=======================================

$ go test -v -bench=. ./...

>> Output:
?       unit-test/entity        [no test files]
Sebelum Unit Test
=== RUN   TestHelloWorldSintya
TestHelloWorld is DONE
--- PASS: TestHelloWorldSintya (0.00s)
=== RUN   TestHelloWorldAjeng
TestHelloWorldLestari is DONE
--- PASS: TestHelloWorldAjeng (0.00s)
=== RUN   TestHelloWorldAssert
TestHelloWorld with Assert is DONE
--- PASS: TestHelloWorldAssert (0.00s)
=== RUN   TestHelloWorld
TestHelloWorld is DONE
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestHelloWorldLestari
TestHelloWorldLestari is DONE
--- PASS: TestHelloWorldLestari (0.00s)
=== RUN   TestHelloWorldRequire
TestHelloWorld with Require is DONE
--- PASS: TestHelloWorldRequire (0.00s)
=== RUN   TestSkip
    skiptest_hello_world_test.go:15: Can't run on Linux
--- SKIP: TestSkip (0.00s)
=== RUN   TestSubTest
=== RUN   TestSubTest/Sintya
=== RUN   TestSubTest/Lestari
--- PASS: TestSubTest (0.00s)
    --- PASS: TestSubTest/Sintya (0.00s)
    --- PASS: TestSubTest/Lestari (0.00s)
=== RUN   TestTableTest
=== RUN   TestTableTest/Sintya
=== RUN   TestTableTest/Lestari
--- PASS: TestTableTest (0.00s)
    --- PASS: TestTableTest/Sintya (0.00s)
    --- PASS: TestTableTest/Lestari (0.00s)
goos: linux
goarch: amd64
pkg: unit-test/helper
cpu: Intel(R) Celeron(R) N5095 @ 2.00GHz
BenchmarkHelloWorld
BenchmarkHelloWorld-4           45079534                24.36 ns/op
BenchmarkHelloWorldLestari
BenchmarkHelloWorldLestari-4    48573410                24.17 ns/op
PASS
Setelah Unit Test
ok      unit-test/helper        2.401s
?       unit-test/repository    [no test files]
=== RUN   TestCategoryService_GetFound
--- PASS: TestCategoryService_GetFound (0.00s)
=== RUN   TestCategoryService_GetSuccess
--- PASS: TestCategoryService_GetSuccess (0.00s)
PASS
ok      unit-test/service       0.008s
*/
