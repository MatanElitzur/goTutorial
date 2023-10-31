// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import "testing"

// https://pkg.go.dev/testing@go1.21.3#T
func TestSum(t *testing.T) { //(t *testing.T) --> This is a pointer of the test runner
	//arrange step
	l, r := 1, 2
	expect := 3
	// action step
	actual := Sum(l, r)
	//assert step
	if expect != actual {
		t.Errorf("Failed to add %v and %v. Actual %v, Expected %v\n", l, r, actual, expect)
	}
}
