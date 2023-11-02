package main // a package is a directory inside a module, a package can be inside of a package.
//All members are visible to other package members
// Example a global scope varibale A in x.go file in main package will be visibale to y.go file in package main
//After download a package with go get command toy can find it at GOPATH='/Users/I044184/go'

//Use the fmt package for printing
//import statment is used to import packages and use their code
//Go will remove the unused packages
import (
	"bufio"
	"demo/knowledge" //Add local (custom) package, demo (the module name) is from the go.mod file
	_ "embed"        //the package is not used in the program but it is imported with the _ (underscore)
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/exp/slices" //This is not part of the standard library so we need to get it with the get command
	//Package alias example
	//"rsc.io/quota"
	//quotev2 "rsc.io/quota/v2"
	//The way we use both versions
	//fmt.Println(quote.Hello())
	//fmt.Println(quotev2.HelloV2())
)

// /////////////////////////////
// //// package members ////////
// ////////////////////////////

// This type struct is a package member so it starts with a capital letter
type User struct {
	ID       int    // ID starts with a capital letter so it considered a public api of the package
	Username string // Username starts with a capital letter so it considered a public api of the package
	password string // password is package level field cause it starts with small letter
	// So if someone will import this User struct he will not see the password field
}

// This function is a package member and it is public cause it starts with a capital letter
func GetById(id int) int {
	fmt.Println(`This function is a package member`)
	return 1
}

// This const is a package member, it is public cause it starts with a capital letter
const MaxUsers = 100

// The init() function is running before the main() function
// Setting global variables, initialize values or XXX establishing a connection
// init() function can be declered on several files at the same time
// and the init() functions will execute on a alphabet file names order
func init() {
	fmt.Println(`*** Hello from the main.go() init() function ***`)
}

// Another example to use embed
// var (
//
//		//go:embed menu.txt
//	    data []slice
//
// )
//

//go:embed menu.txt
var sread string

func main() {
	fmt.Println("Used 'embed' to read content from file and set to variable sread")
	fmt.Println(sread)
	fmt.Println("*************************************************")
	//println("Hello, Gophers!")
	fmt.Println(`This is a Raw string - Hello, Gophers! \n Hello`)
	fmt.Println("This is interpreted string - Hello, Gophers! \n Hello")
	//////////////////////////////////////
	/////// Variables example ////////////
	//////////////////////////////////////
	fmt.Println(`*****  VARIABLES  ******`)
	var myName string //Declare variable
	myName = "Mike"
	fmt.Println(myName)
	var myName1 string = "Mike" //Declare and initialize variable
	fmt.Println(myName1)
	var myName2 = "Mike" //Initialize with variable with inferred type
	fmt.Println(myName2)
	myName3 := "Mike" //short declaration and initialize syntax
	fmt.Println(myName3)
	var i int = 45
	var f float32
	f = float32(i) //Type conversions allow explicit conversion
	fmt.Println(f)
	a, b := 10, 5 //Go allows multiple variables to be initialized at once!
	c := a + b
	fmt.Println(c)
	d := a == b
	fmt.Println(d) //false
	e := a != b
	fmt.Println(e) //true
	g := a <= b
	fmt.Println(g) //false
	//////////////////////////////////////
	/////// Constants example ////////////
	//////////////////////////////////////
	fmt.Println(`*****  CONSTANTS  ******`)
	const con = 42                   // constant(Implicitly typed)
	const bb string = "hello, world" //Explicitly typed constant
	const cc = con                   //one constant can be assigned to another
	const (                          //group of constants
		dd = true
		ee = 3.14
		ii = 2 * 5
		ss = `Hello`
	)

	const (
		a2 = 8
		b2 = 5
		c2 //c2 will also get the value 5
	)
	const ( //iota is the position in the const
		b1 = iota     //0
		c1            //1
		d1 = 3 * iota //6
	)

	//////////////////////////////////////
	/////// Pointers example ////////////
	//////////////////////////////////////
	fmt.Println(`*****  POINTERS  ******`)
	a_var := 42
	fmt.Println(a_var)
	b_var := &a_var    //b_var is a pointer he points to a_var address ;
	fmt.Println(b_var) //prints the memory address
	//*b_var (dereferencing the pointer value) result 42
	fmt.Println(*b_var) //42
	*b_var = 1000
	fmt.Println(a_var)  //1000
	b_var = new(int)    //buit-in "new" function creates pointer to anonymous variable
	fmt.Println(b_var)  //The b_var memory address is changed
	fmt.Println(*b_var) //The b_var pointer value is equel to zero cause it was not assigned

	/////////////////////////////////////////////////
	/////// Aggregate Data Types example ////////////
	////////////////////////////////////////////////
	//Arrays has a determind size it can't grow or shrink
	fmt.Println(`***** AGGREGATE DATA TYPES ******`)
	fmt.Println(`***** AGGREGATE DATA TYPES --> ARRAY ******`)
	var arr [3]int //Declare array of type int with 3 places
	fmt.Println(arr)
	arr = [3]int{1, 2, 3} //Initialize an array
	fmt.Println(arr)
	fmt.Println(arr[1]) //Print a single element of an array
	arr[1] = 99         //Update a value
	fmt.Println(arr)
	fmt.Println(len(arr)) //print the length of the array
	arr_str := [3]string{"foo", "bar", "baz"}
	fmt.Println(arr_str)
	arr_str2 := arr_str //This will do a copy operation (It is not referring to the same data in the memory)
	fmt.Println(arr_str2)
	arr_str[0] = "quux"
	fmt.Println(arr_str)             //The array will have quux in the first location in the array
	fmt.Println(arr_str2)            // The values in this array are the same as before without quux in the first location
	fmt.Println(arr_str == arr_str2) //compare between arrays, this will return false cause there is a different with the values, It's compare each value one at a time.

	fmt.Println(`***** AGGREGATE DATA TYPES --> SLICES ******`)
	//Slices has the ability to grow and shrink
	//Slices do not storing their own data they refer to a data that is been stored by an array
	var s1 []int           //Slices of ints
	fmt.Println(s1)        // [] (nil)
	fmt.Println(s1 == nil) // return true cause we are not pointing
	s1 = []int{1, 2, 3}    //Initialize a slice
	fmt.Println(s1[1])
	s1[1] = 99 //Update a value in a slice
	fmt.Println(s1)
	s1 = append(s1, 5, 10, 15) // add elements to the slice
	fmt.Println(s1)
	s1 = slices.Delete(s1, 1, 3) //Remove indices 1, 2 from slice (golang.org/x/exp/slices), Why 3 --> cause it until 3, not include 3
	fmt.Println(s1)
	s2 := s1
	s2[2] = 777 //When we changed the value, both slices values changed as well
	fmt.Println(s1, s2)
	s1 = append(s1, 100, 200) //The 100 & 200 will be added only to s1 and not to s2
	fmt.Println(s1, s2)

	fmt.Println(`***** AGGREGATE DATA TYPES --> Maps ******`)
	var m1 map[string]int //Declare a map that has a keys of type string and values of type int
	fmt.Println(m1)
	fmt.Println(m1 == nil)                  // return true
	m1 = map[string]int{"foo": 1, "bar": 2} //Initialize a map
	fmt.Println(m1)
	fmt.Println(m1["foo"]) //Lookup a value in map
	m1["bar"] = 99         //Update a value in map
	delete(m1, "foo")      //Remove entry from map
	fmt.Println(m1)
	fmt.Println(m1["foo"]) //0 - queries always return results although we deleted foo key
	v, ok := m1["foo"]     //To solve the issue in the line before we are using the following
	fmt.Println(v)         // Still return 0
	fmt.Println(ok)        // false - The value not exists in the map cause the key was not present in the map
	m2 := m1               // maps are copied by reference
	// use maps.Clone to clone maps without the references
	m1["foo"], m2["bar"] = 99, 42 //Update values in map
	fmt.Println(m1, m2)           //As we can see the data is shared
	//fmt.Println(m1 == m2)         // We can't compare 2 maps, only a mpa to a nil
	var m3 map[string][]string //Map, key of string and the value is a slice of string
	fmt.Println(m3)
	m3 = map[string][]string{ //Init map with string slice values
		"coffee": {"Coffee", "Espresso", "Cappccino"},
		"tea":    {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
	fmt.Println(m3)
	fmt.Println(m3["coffee"])
	m3["other"] = []string{"Hot Chocolate"} // Add a new slice into the map of string slices
	fmt.Println(m3)
	delete(m3, "tea")
	fmt.Println(m3)
	fmt.Println(m3["tea"]) //tea not exists so we will get an uninitialize slice
	v3, ok3 := m3["tea"]
	fmt.Println(v3)  // Still return 0
	fmt.Println(ok3) // false - The value not exists in the map cause the key was not present in the map
	m4 := m3
	m4["coffee"] = []string{"Coffee"}
	m3["tea"] = []string{"Hot Tea"}
	fmt.Println(m3)
	fmt.Println(m4)
	//They are both equel pointing to the same under line data structure

	fmt.Println(`***** AGGREGATE DATA TYPES --> Structs ******`)
	var stru struct { //Declare an anonymous struct
		name string
		id   int
	}
	fmt.Println(stru)
	stru.name = "Matan"
	fmt.Println(stru.name)

	type myStruct struct { //Create custom type based on struct
		name string
		id   int
	}
	var s4 myStruct //Declare variable with custom type
	fmt.Println(s4)
	s4 = myStruct{ //Init the struct (struct literal)
		name: "Shaili",
		id:   42}
	fmt.Println(s4)
	s5 := s4 //We are doing a copy! there is no reference to the same values in the memory
	s4.name = "Mike"
	fmt.Println(s4, s5) //They are different, cause struct are value types like arrays

	// type menuItem struct {
	// 	name   string
	// 	prices map[string]float64 //map the keys are string the values are float64
	// }

	// menu := []menuItem{ //Create a slice of type menuItem, menueItem is a struct
	// 	{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
	// 	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	// }

	//fmt.Println(menu)

	//////////////////////////////////////
	/////// Use Loops ////////////
	//////////////////////////////////////
	fmt.Println(`*****  LOOPS  ******`)
	knowledge.InfiniteLoop()
	knowledge.LoopTillCondition()
	knowledge.CounterBasedLoop()
	knowledge.LoopingWithCollections()
	//loopingWithCollections1(menu) //menu was declered inside the main function
	knowledge.LoopingWithCollections2() //menu is declered in the flobal scope
	//instance := knowledge.Menu{}        //We crated an instance of type Menu
	//instance.LoopingWithCollections3()
	//////////////////////////////////////
	/////// Use Branching ////////////
	//////////////////////////////////////
	fmt.Println(`*****  BRANCHING  ******`)
	knowledge.If_expression()
	knowledge.If_expression1()
	knowledge.Switch_expression()
	knowledge.Switch_expression1()
	knowledge.DeferredFunc1()
	knowledge.PanicAndRecover()
	knowledge.PanicAndRecoverDivideExample()
	//////////////////////////////////////
	/////// Use functions ////////////
	//////////////////////////////////////
	fmt.Println(`*****  FUNCTIONS  ******`)

	knowledge.Greet("Matan", "Mike")
	knowledge.Greet1("Oren", "Golan")
	name := "Kfir"
	name1 := "Max"
	name2 := "Irena"
	name3 := "Roy"
	name4 := "Yaniv"
	knowledge.Greet2(name, name1, name2, name3, name4)
	knowledge.Greet3(name, &name1)
	result1 := knowledge.Sum(3, 4)
	fmt.Println(result1)
	result2, result2Bool := knowledge.DivideNumber(1, 5)
	fmt.Println(result2, result2Bool)
	result3, result3Bool := knowledge.DivideNumber1(1, 5)
	fmt.Println(result3, result3Bool)

	//////////////////////////////////////
	/////// Use Interfaces ///////////////
	//////////////////////////////////////
	fmt.Println(`*****  INTERFACES  ******`)
	var p knowledge.Printer                        // creating a variable of type Printer from module knowledge
	p = knowledge.User{Username: "Popsi", Id: 123} //assign new object of type knowledge.User to the knowledge.Printer variable and initialize it
	fmt.Println(p.Print())
	p = knowledge.Item{Itemname: "Hammer", Id: 1}
	fmt.Println(p.Print())

	//Create go Panic
	// testInterface := p.(knowledge.User) // I expect that p is of type knowledge.User, if I am wrong go will panic
	// testInterface.Print()

	//Now I will not get Go Panic
	testInterface, ok := p.(knowledge.User) // I expect that p is of type knowledge.User, if I am wrong go will panic
	fmt.Println(testInterface, ok)          //// Now we get false cause the assertion is wrong

	testInterface1, ok := p.(knowledge.Item)
	fmt.Println(testInterface1, ok) // Now we get true cause the assertion is right
	switch v := p.(type) {
	case knowledge.User:
		fmt.Println("Found a user!", v)
	case knowledge.Item:
		fmt.Println("Found an item!", v)
	default:
		fmt.Println("I am not sure what is is...", v)
	}

	//////////////////////////////////////
	/////// Use Interfaces ///////////////
	//////////////////////////////////////
	fmt.Println(`*****  GENERIC  ******`)
	knowledge.Clone()

	////////////////////////////////////////////
	/////// Error Managment ////////////////////
	////////////////////////////////////////////
	fmt.Println(`*****  ERROR MANAGMENT  ******`)
	knowledge.ErrorManagment1()

	////////////////////////////////////////////
	/////// Error Managment ////////////////////
	////////////////////////////////////////////
	fmt.Println(`*****  CONCURRENT PROGRAMING  ******`)
	knowledge.ConcurrentPrograming1()

	//////////////////////////////////////
	/////// WebServer ////////////
	//////////////////////////////////////
	// fmt.Println(`*****  WebServer  ******`)
	// http.HandleFunc("/", Handler)
	// http.ListenAndServe(":3000", nil) //"localhost:3000"
	// //http://localhost:3000/

	//////////////////////////////////////
	/////// Use console ////////////
	//////////////////////////////////////
	fmt.Println(`*****  CONSOLE  ******`)
	fmt.Println(`What would you like me to scream?`)
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n') //The function ReadString return a string and an error but we will not manage the error so we will type _ and ignore the error
	s = strings.TrimSpace(s)    //Will remove the new line (\n)
	s = strings.ToUpper(s)
	fmt.Println(s + `!`)

loop: //This is a label
	for {
		fmt.Println(`Please select an option`)
		fmt.Println(`1) Print Menu`)
		fmt.Println(`2) Add item`)
		fmt.Println(`q) quite`)
		in_01 := bufio.NewReader(os.Stdin)
		choice, _ := in_01.ReadString('\n') //The function ReadString return a string and an error but we will not manage the error so we will type _ and ignore the error
		switch strings.TrimSpace(choice) {
		case "1":
			//loopingWithCollections1(menu)//menu was declered in the main function
			knowledge.LoopingWithCollections2() //menu is now in the global scope
		case "2":
			fmt.Println(`Please enter the name of the new item`)
			name, _ := in_01.ReadString('\n')
			//add a new item in the menu and create an empty map with the make command
			addItemError := knowledge.AddItem(name)
			if addItemError != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", addItemError))
			}
		case "q":
			break loop //It will break the stetment after the label (The infinite for loop)
		default:
			fmt.Println(`Unknown option`)
		}

	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt") // _ we ignore the error
	io.Copy(w, f)                 //We copy the content of the file to the response
}
