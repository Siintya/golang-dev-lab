package helper

import (
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Sintya")

	// Jika pengecekan gagal, akan memanggil Fail()
	assert.Equal(t, "Hello Sintya", result, "Result must be 'Hello Sintya'")
	/* Output jika error:
	=== RUN   TestHelloWorldAssert
			assertion_hello_world_test.go:14:
						Error Trace:    /media/sweetcorn/AS-SSD/Backup 31jul2025/2026/Documents/Lesson & Learn/Programming-Lang/Golang (GO)/golang-dev-lab/04-unit-test/helper/assertion_hello_world_test.go:14
						Error:          Not equal:
										expected: "Hello Sintya"
										actual  : "HI Sintya"

										Diff:
										--- Expected
										+++ Actual
										@@ -1 +1 @@
										-Hello Sintya
										+HI Sintya
						Test:           TestHelloWorldAssert
						Messages:       Result must be 'Hello Sintya'
		TestHelloWorld with Assert is DONE
		--- FAIL: TestHelloWorldAssert (0.00s)
		FAIL
		exit status 1
		FAIL    04-unit-test/helper     0.006s
	*/

	fmt.Println("TestHelloWorld with Assert is DONE")
}
