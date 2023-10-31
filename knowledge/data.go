// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import "fmt"

// The init() function is running before the main() function
// Setting global variables, initialize values or XXX establishing a connection
// init() function can be declered on several files at the same time
// and the init() functions will execute on a alphabet file names order
func init() {
	fmt.Println(`*** Hello from the data.go() init() function ***`)
}

// ////////////////////////////////////
// //// Public package members ////////
// ///////////////////////////////////
// This type struct is a package member so it starts with a capital letter
type MenuItem struct {
	name   string             // name is private cause it starts with a small letter
	prices map[string]float64 //map the keys are string the values are float64 ;  name is private, it starts with a small letter
}

// This is an object orinthed flow, we use this way cause we want to add methods for type menu
// see function func (m menu) LoopingWithCollections3()
type Menu []MenuItem //The underline type is a slice collection of MenuItem

// This variable is a package member
// var data = []MenuItem{ //Create a slice of type menuItem, menueItem is a struct
//
//		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
//		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
//	}

// we created an instance of Menu and initialized it
var data = Menu{ //Create a slice of type menuItem, menueItem is a struct
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}
