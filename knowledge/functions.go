// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	"errors"
	"fmt"
	"strings"
)

// Public package function --> start with a capital letter
func Greet(name1 string, name2 string) {
	fmt.Println(name1)
	fmt.Println(name2)
}

// Public package function --> start with a capital letter
func Greet1(name1, name2 string) { //name1 and name2 are both defined as string
	fmt.Println(name1)
	fmt.Println(name2)
}

// Public package function --> start with a capital letter
func Greet2(names ...string) { //Variadic Parameters --> Collaction of strings , Variadic parameters must be final parameter
	for _, n := range names {
		fmt.Println(n)
	}
}

// Public package function --> start with a capital letter
// Use pointers to share memory otherwise use values
func Greet3(name1 string, name2 *string) { //--> name2 is a pointer
	fmt.Println(name1)
	fmt.Println(*name2) // print the value
	fmt.Println(name2)  //prints the address
}

// Public package function --> start with a capital letter
func Sum(a, b int) int {
	return a + b
}

// Public package function --> start with a capital letter
func DivideNumber(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// Public package function --> start with a capital letter
func DivideNumber1(a, b int) (result int, ok bool) { //rarley used
	if b == 0 {
		return //This will reurn --> 0, false
	}
	result = a / b
	ok = true
	return
}

func AddItem(item string) error {
	item = strings.TrimSpace(item)
	for _, itemValue := range data {
		if itemValue.name == item {
			return errors.New("menu item already exists")
		}
	}
	data = append(data, MenuItem{name: item, prices: make(map[string]float64)})
	return nil //We must to return a value cause this value return an error
	//In go nil indicates that now error was encountered
}
