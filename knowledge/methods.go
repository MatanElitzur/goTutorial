// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import "fmt"

var i int

// ///////////////////////
// This is a function ///
// ///////////////////////
func isEvenFunc(i int) bool {
	return i%2 == 0
}

var ans = isEvenFunc(i) // This is how we call the upper function

// /////////////////////////
// //  This is a method ////
// ////////////////////////
// We can declare a type of anything
type myInt int // Declare a type by the name myInt and it is of type int
var mi myInt   // declare a variable of type myInt
// This is a method; (i myInt) this part call a method receiver
// The (i myInt) is doing a tight coupling between the type myInt and the isEven function
// when there is a tight coupling between a type and a func, the func become a method
// The method will excecute always in a context of a variable in the following case the variable i
func (i myInt) isEven() bool {
	return int(i)%2 == 0
}

var ans1 = mi.isEven()

type user struct {
	id       int
	username string
}

// Here we pass user as a value
func (u user) String() string { // value receiver
	return fmt.Sprintf("%v (%v)\n", u.username, u.id)
}

// Here we pass user as a pointer so we share a variable between a caller and a method
// func (u *user) UpdateName(n name) { //pointer receiver
// 	u.username = name
// }

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// The use of (sv SemanticVersion) meanes a value base receiver it creates a copy of the object that is been sent
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// The use of (sv *SemanticVersion) meanes a reference base receiver it works on the object that invoked this func
func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

// The use of (sv *SemanticVersion) meanes a reference base receiver it works on the object that invoked this func
func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}

// The use of (sv *SemanticVersion) meanes a reference base receiver it works on the object that invoked this func
func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}
