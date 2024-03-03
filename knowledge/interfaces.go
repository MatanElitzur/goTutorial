// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import "fmt"

//Interface allow me to describe the behavior interface require
// type Reader interface { //The name of the interface is Reader
// 	Read([]byte) (int, error) //Read method that takes a slice of byte and return an in and an error
// }

// type File struct {}

// func (f File) Read(b []byte) (n int, err error) { // The read method looks like the Read method in the interface

// }

// // Create another struct that will have the same Read method like in the interface
// type TCPConn struct {}

// func (t TCPConn) Read(b []byte) (n int, err error) { // The read method looks like the Read method in the interface

// }

// var f File    //Create a variable of the struct  type File
// var t TCPConn //Create a variable of the struct type TCPConn
// var r Reader  //Create a variable of the interface type
// r = f //Allow us to assing the f struct variable to the r interface variable
// r.Read() //We can call the Read method from the File struct
// r = t //Allow us to assing the t struct variable to the r interface variable
// r.Read() //We can call the Read method from the TCPConn struct

// var r1 Reader = f  //Create a variable of the interface type and assign f variable
// var f2 File = r1 //Error, go can't be sure this will work
// f2 = r1.(File) //type assertion, panic uppon failure
// f2, ok := r1.(File) //type assertion, with comma okay, doesn't panic

// //TYPE SWITCHES
// switch v := r.(type) {
// case File:
// 	//v now is a File object
// case TCPConn:
// 	//v now is a TCPConn object
// default:
// 	//this is selceted if no types were matched
// }

type Printer interface {
	Print() string
}

type User struct {
	Username string
	Id       int
}

type Item struct {
	Itemname string
	Id       int
}

// Implement the interface for User struct
func (u *User) Print() string {
	fmt.Println(`interfaces.go --> Inside Print method of struct User`)
	return fmt.Sprintf("%v [%v]\n", u.Username, u.Id)
}

// Add method to the User struct, this method not exists in the interface
func (u *User) Hello() string {
	fmt.Println(`interfaces.go --> Inside Hello method of struct User`)
	return fmt.Sprintf("Hello %v [%v] how are you?\n", u.Username, u.Id)
}

// **********************************************************
// Implement the interface for Item struct
func (i *Item) Print() string {
	fmt.Println(`interfaces.go --> Inside Print method of struct Item`)
	return fmt.Sprintf("%v [%v]\n", i.Itemname, i.Id)
}

// Add method to the Item struct, this method not exists in the interface
func (i *Item) Color() string {
	fmt.Println(`interfaces.go --> Inside Color method of struct Item`)
	return fmt.Sprintf("Color method, %v [%v]\n", i.Itemname, i.Id)
}
