// Package name and folder name must be identical for the go language to recognize the package
// Package knowledge has an examples of golang code
package knowledge

import (
	"fmt"
)

// Public package function --> start with a capital letter
func If_expression() {
	var test bool //default is false
	var test1 bool = true
	if test {
		fmt.Println("test is true")
	}

	if test {
		fmt.Println("test is true")
	} else {
		fmt.Println("else at the end")
	}

	if test {
		fmt.Println("test is true")
	} else if test1 {
		fmt.Println("test1 is true")
	} else {
		fmt.Println("else at the end")
	}

}

// Public package function --> start with a capital letter
func If_expression1() {
	if i := 15; i < 15 {
		fmt.Println("i is less then 15")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("after the if statement")

}

// Public package function --> start with a capital letter
func Switch_expression() {
	i := 5
	switch i {
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*i + 3:
		fmt.Println("secound case")
	default:
		fmt.Println("default case")
	}
}

// Public package function --> start with a capital letter
func Switch_expression1() {
	//Logical switch
	switch i := 8; { //true is implied
	case i < 5:
		fmt.Println("i is less then 5")
	case i < 10:
		fmt.Println("i is less then 10")
	default:
		fmt.Println("i is greater then 10")
	}
}

// Public package function --> start with a capital letter
func DeferredFunc1() {
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")
}

// Public package function --> start with a capital letter
func PanicAndRecover() {
	fmt.Println("panicAndRecover main 1")
	panicAndRecover1()
	fmt.Println("panicAndRecover main 2")
}

func panicAndRecover1() {
	defer func() { //Anonymous function
		msg := recover() //recover saves the calling function to be terminated
		fmt.Println(msg)
	}()
	fmt.Println("func panicAndRecover1 #1")
	panic("uh-oh")
	fmt.Println("func panicAndRecover1 #2")
}

// Public package function --> start with a capital letter
func PanicAndRecoverDivideExample() {
	fmt.Println("panicAndRecoverNext main 1")
	divide(10, 0)
	fmt.Println("panicAndRecoverNext main 2")
}

func divide(dividend int, divisor int) int {
	defer func() { //Anonymous function
		//get the result to msg and then check if msq is not nil (Not nil mean app is panic)
		//cause at the first time when we divide numbers that can be divided we get nil
		//recover() is part of the recaver panic mechanism,
		// the recover will recover the application if it panicking
		if msg := recover(); msg != nil { //recover saves the calling function to be terminated
			//a panic has been detected in the application
			fmt.Println(msg)
		}
	}()
	fmt.Println("func panicAndRecover2 #1")
	sum := dividend / divisor
	fmt.Println("func panicAndRecover2 #2")
	return sum
}
