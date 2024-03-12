// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ErrorManagment1() {
	error1()
	error2_createAnError()
	resultDivideNumber2, errDivideNumber2 := divideNumber2(10, 0)
	if errDivideNumber2 != nil {
		fmt.Println(errDivideNumber2)
	}
	fmt.Println(resultDivideNumber2)

	resultDivideNumber3, errDivideNumber3 := divideNumber3(10, 0)
	if errDivideNumber3 != nil {
		fmt.Println(errDivideNumber3)
	}
	fmt.Println(resultDivideNumber3)

	errorExample1()
	panicRecover1()
	repanicking()
}

// processInput simulates processing of user input
func processInput(input string) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic for debugging purposes
			log.Printf("Recovered in processInput: %v", r)
			// Perform any necessary cleanup
			fmt.Println("Cleanup after panic...")
			// Re-panic to ensure the calling function
			// can handle it or fail gracefully
			panic(r)
		}
	}()

	// Simulate a panic scenario based on input
	// If the input is not a number, panic
	if _, err := strconv.ParseInt(input, 10, 64); err != nil {
		panic(err)
	}

	fmt.Println("Processed input successfully:", input)
}

func repanicking() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Main recovered from panic:", r)
		}
	}()

	// Get a number from the user
	fmt.Print("Enter a number: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read from standard input:", err)
		return
	}
	input = strings.Trim(input, "\n")

	fmt.Println("Processing input:", input)
	processInput(input)

	fmt.Println("Main function continues after processing input.")
}

func panicRecover1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panicRecover1 Panic recovered: %v\n", r)
		}
	}()
	panic("panicRecover1 Random panic!!!")
}

func error1() {
	f, err := os.Open("path/to/file")
	if err != nil {
		fmt.Println("knowledge-->error1 ", err)
	}
	defer f.Close()
}

func error2_createAnError() {
	err := errors.New("knowledge-->error2 This is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("this error wraps the first error: %w", err)
	fmt.Println(err2)
}

func divideNumber2(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return l / r, nil
}

// func divideNumber3(l, r int) (int, error) {
func divideNumber3(l, r int) (result int, err error) {
	defer func() { //Anonymous function
		//get the result to msg and then check if msq is not nil (Not nil mean app is panic)
		//cause at the first time when we divide numbers that can be divided we get nil
		//recover() is part of the recaver panic mechanism,
		// the recover will recover the application if it panicking
		if msg := recover(); msg != nil { //recover saves the calling function to be terminated
			//a panic has been detected in the application
			//fmt.Println(msg)
			fmt.Printf("knowledge-->errorManagment-->divideNumber3 %v", msg)
			result = 0
			err = fmt.Errorf("%v", msg)
			//Example1 --> err = errors.New("invalid divisor: must not be zero")
			//This is what we want --> return 0, errors.New("invalid divisor: must not be zero") --> This will not work we need to add names to the function return values
		}
	}()
	return l / r, nil
}

// /////////////////////////////////////////////////////////////////////////////////////////////
func errorExample1() error {
	var r io.Reader = BadReader{err: errors.New("my nonsense reader")}
	// if _, err := r.Read([]byte("test something")); err != nil {
	// 	fmt.Printf("an error occurred %s", err)
	// 	return err
	// }
	value, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("an error occurred %s", err)
		return err
	}
	fmt.Println(value)
	return nil
}

// BadReader struct is implementing the inferface Reader
type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}
