// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	//The crypto/sha1, crypto/sha256, and crypto/sha512 packages are imported for hashing operations.
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

//import "testing/quick" --> for black box testing, if we don't care about the internal of a package
//import "testing/iotest" --> test package for readers and writeters data from an external source
//import "net/http/httptest" --> simulate request and response and set test servers that we can use for testing
//Useful community projects
//1. Testify  github.com/stretchrcom/testify  More assertion options
//2. Ginkgo   github.com/onsi/ginkgo bdb still --> behavior-driven style
//3. GoConvey goconvey.co --> how to view the test results in a browser base format
//4. httpexpect github.com/gavv/httpexpect --> Good for Rest test api
//5. gomock code.google.com/p/gomock/ --> allow you to mock objects
//6. go-sqlmock github.com/DATA-DOG/go-sqlmock -->

// Documentation
// https://golang.org/pkg/testing
// https://pkg.go.dev/testing@go1.21.3#T

//Non-immediate failures
//t.Fail()
//t.Error(...interface{})
//t.Errorf(string, ...interface{})

// Immediate failures
// t.FailNow()
// t.Fatal(..interface)
// t.Fatalf(string, ..interface)

//Run tests in Parallel
// t.Parallel --> You can add t.Parallel at the begining of the tests and the tests will run in Parallel
//You must verify that the tests are not using the same resources

//Skip a test
//t.Skip() // You can add this line at the begining of the test and the test will skip

func TestSum(t *testing.T) { //(t *testing.T) --> This is a pointer of the test runner
	t.Parallel()
	//arrange step
	t.Log("Sarted TestSum func") // Will print to the log only if test fails
	l, r := 1, 2
	t.Logf("Sum func input are %d and %d", l, r) // Will print to the log only if test fails
	expect := 3
	// action step
	actual := Sum(l, r)
	//assert step
	if expect != actual {
		t.Errorf("Failed to add %v and %v. Actual %v, Expected %v\n", l, r, actual, expect)
	}
}

// This way of writing test is powerful cause you can test different options as inputs
func TestSumTableDriven(t *testing.T) {
	t.Parallel()
	scenarios := []struct {
		inputA int
		inputB int
		expect int
	}{
		{inputA: 1, inputB: 2, expect: 3},
		{inputA: 2, inputB: 2, expect: 4},
	}

	for _, s := range scenarios {
		got := Sum(s.inputA, s.inputB)
		if got != s.expect {
			t.Errorf("Did not get expected result inputA '%v', inputB '%v'. Expected '%v', got '%v'", s.inputA, s.inputB, s.expect, got)
		}
	}
}

//Immediate failure
//(t *testing.T) --> This is a pointer of the test runner
//t.FailNow()
//t.Fatal(ars ...interface{})
//t.Fatalf(format string, ars ...interface{})

//Non-Immediate failure
//(t *testing.T) --> This is a pointer of the test runner
//t.Fail()
//t.Error(ars ...interface{})
//t.Errorf(format string, ars ...interface{})

//More functions
//Skip --> allow us to skip a test
//Skipf --> allow us to skip a test with printing a string
//SkipNow -->
//Run --> use a callback in a test functions
//Parallel --> Option to run the tests in parallel

// //////////////////////////////////////
// ///////// Performance tests /////////
// /////////////////////////////////////
// Benchmark tests
func BenchmarkSum(b *testing.B) {
	//b.StartTimer
	//b.StopTimer
	//b.RestTimer
	//b.RunParallel()
	//arrange step
	l, r := 1, 2
	// action step
	for i := 0; i < b.N; i++ { //The b.N - is how much iteration the test will run
		Sum(l, r)
	}
}

func BenchmarkSHA1(b *testing.B) {
	//BenchmarkSHA1, BenchmarkSHA256, and BenchmarkSHA512: Benchmark the performance of SHA1, SHA256, and SHA512 hashing functions, respectively.
	data := []byte("Mary has a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ { //The b.N value is a measure of how many times the test should be run. The higher the b.N value, the more accurate the benchmark will be.
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	//BenchmarkSHA1, BenchmarkSHA256, and BenchmarkSHA512: Benchmark the performance of SHA1, SHA256, and SHA512 hashing functions, respectively.
	data := []byte("Mary has a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ { //The b.N value is a measure of how many times the test should be run. The higher the b.N value, the more accurate the benchmark will be.
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	//BenchmarkSHA1, BenchmarkSHA256, and BenchmarkSHA512: Benchmark the performance of SHA1, SHA256, and SHA512 hashing functions, respectively.
	data := []byte("Mary has a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ { //The b.N value is a measure of how many times the test should be run. The higher the b.N value, the more accurate the benchmark will be.
		sha512.Sum512(data)
	}
}
