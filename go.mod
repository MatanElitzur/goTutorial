//The go.mod is required in the root of the project
//It defined the module and it dependencies in the go modules
module demo // The current name of the module

go 1.23.0 // The go version the module is compatible

require golang.org/x/exp v0.0.0-20231006140011-7918f672742d

require (
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/time v0.5.0 // indirect
)

//require github.com/pioz/faker v1.7.3
