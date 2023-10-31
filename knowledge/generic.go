// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import "fmt"

// Generic are better then iterface casue when we assign a concreate object to an interface we lose his identity
// Example:
// interface := concreate
// interface. --> will show only the methods of the interface
var testScores64 = []float64{
	87.3,
	105,
	63.5,
	27,
}

var testScores32 = []float32{
	8.3,
	10,
	6.5,
	2,
}

var testScoresMap = map[string]float64{
	"Harry":    8.3,
	"Hermione": 10,
	"Ronald":   6.5,
	"Neville":  2,
}

var a1 = []int{1, 2, 3}
var a2 = []float64{3.14, 6.02}
var a3 = []string{"foo", "bar", "baz"}

//c := testScores //This will cause a reference copy cause the type is slice

func Clone() {
	c64 := clone(testScores64)
	fmt.Println(&testScores64[0], &c64[0], c64)

	c32 := clone(testScores32)
	fmt.Println(&testScores32[0], &c32[0], c32)
	// Using map for generic
	map64 := cloneMap(testScoresMap)
	fmt.Println(map64)

	///Using an interface for generic
	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)
	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

// Clone the incoming slice
// func clone(s []float64) []float64 {
// 	result := make([]float64, len(s)) // create a slice with the same lenght as the income slice
// 	for i, v := range s {
// 		result[i] = v
// 	}
// 	return result
// }

// Clone the incoming slice
func clone[V any](s []V) []V {
	fmt.Println("knowledge-->clone")
	result := make([]V, len(s)) // create a slice with the same lenght as the income slice
	for i, v := range s {
		result[i] = v
	}
	return result
}

//	func cloneMap(m map[string]float64) map[string]float64 {
//		result := make(map[string]float64, len(m)) // create a map with the same lenght as the income slice
//		for k, v := range m {
//			result[k] = v
//		}
//		return result
//	}
func cloneMap[K, V comparable](m map[K]V) map[K]V {
	fmt.Println("knowledge-->cloneMap")
	result := make(map[K]V, len(m)) // create a map with the same lenght as the income slice
	for k, v := range m {
		result[k] = v
	}
	return result
}

/////////////////////////////////////////////////

type addable interface {
	int | float64 | string //Allow the following types
}

func add[V addable](s []V) V {
	fmt.Println("knowledge-->add")
	var result V
	for _, v := range s { //The _ is for the index, but we don't need it and therefor it is _
		result += v
	}
	return result
}
