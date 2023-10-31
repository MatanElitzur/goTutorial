// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	"fmt"
	"strings"
)

// Public package function --> start with a capital letter
func InfiniteLoop() {
	i := 1
	for {
		fmt.Println(i)
		i += 1
		break //We added break to exit the infinte loop
	}
}

// Public package function --> start with a capital letter
func LoopTillCondition() {
	i := 1
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("Done!")
}

// Public package function --> start with a capital letter
func CounterBasedLoop() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("Done!")
}

// Public package function --> start with a capital letter
func LoopingWithCollections() {
	arr := [3]int{101, 102, 103}
	for i, v := range arr { // i == index, v == value
		fmt.Println(i, v)
	}
	fmt.Println("Done!")

	for _, v := range arr { // _ means we don't want to use the index, v == value
		fmt.Println(v)
	}
	fmt.Println("Done!")
}

// Public package function --> start with a capital letter
func LoopingWithCollections2() { // data is in a global scope therefore we can loop
	for _, item := range data { // _ mean ignore the index, v == value
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			//Printf --> formatted print statement
			fmt.Printf("\t%10s%10.2f\n", size, price) //\t --> shift tab ; %10s --> 10 string characters ; %10.2 --> 10 characters column with 2 decimal precision ; f --> expect float number
		}
	}
	fmt.Println("Done!")
}

// Public package function --> start with a capital letter
func LoopingWithCollections3(slice []MenuItem) { // slice is been copied when used from the caller func
	for _, item := range slice { // _ mean ignore the index, v == value
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			//Printf --> formatted print statement
			fmt.Printf("\t%10s%10.2f\n", size, price) //\t --> shift tab ; %10s --> 10 string characters ; %10.2 --> 10 characters column with 2 decimal precision ; f --> expect float number
		}
	}
	fmt.Println("Done!")
}

// (m menu) this is the method reciver we are passing here --> basically this method belongs to type menu
// Public package function --> start with a capital letter
func (m Menu) LoopingWithCollections3() { // slice is been copied when used from the caller func
	fmt.Println("start LoopingWithCollections3 --> *** This Method is belons to type Menu ***")
	//for _, item := range m.getData() { // _ mean ignore the index, v == value
	for _, item := range data { // _ mean ignore the index, v == value
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			//Printf --> formatted print statement
			fmt.Printf("\t%10s%10.2f\n", size, price) //\t --> shift tab ; %10s --> 10 string characters ; %10.2 --> 10 characters column with 2 decimal precision ; f --> expect float number
		}
	}
	fmt.Println("Done!")
	fmt.Println("end LoopingWithCollections3 --> *** This Method is belons to type Menu ***")
}

// This is a private method of Menu cause it start with lower case letter
func (m Menu) getData() Menu {
	return data
}
