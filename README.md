# goTutorial
## useful links
1. https://go.dev/
2. https://pkg.go.dev/std --> Standard library
3. https://go-proverbs.github.io/ --> Do & Don't
4. testing package --> https://pkg.go.dev/testing@go1.21.3#T
## go commands
1. go mod --> module maintenance, will display several commands
2. go mod init <NAME OF MODULE>--> initialize a new module in current directory, create a go.mod file (it's a configuration file about the module and it dependencies) inside the module name is demo.
   1. Example: go mod init ps.m3.demo1 
   2. Example: go mod init demo
3. go build --> compile the code and create a binary file in the directory.
4. go run . --> compile the code into a temp directory and run the Go program.(The dot is for the current directory)
5. go get <package> --> Thie will download a package
   1. Example: go get golang.org/x/exp/slices
   2. Example: go get github.com/pioz/faker
   3. Example: download a specific version - go get github.com/pioz/faker@v1.6.0
   4. Example go get -u --> Automatically updated the packages to latest version
6. go test <Tests locations> 
   1.  go test . --> Run all tests in root directory
   2.  go test ./knowledge --> Run all tests in knowledge directory
   3.  got test ./...  --> Run all tests in root directory and subsirectories (All test in the module)
7. go mod tidy --> remove unused dependencies from go.mod file.
8. go mod vendor -v --> Create a vendor folder in the same folder(root folder), it will contain the packages that we are using
9. go env --> display all go variables, 
   1. Example GOPATH variable where packages are located
   2. Example GOMODCACHE where the packages are stored in the cache
   3. go parameter GO111MODULE --> Used when we want to use am external tool to manage modules dependencies
      1. go env -w GO111MODULE=off --> Remove the Go module awareness means go command looks for packages in the directories specified by the GOPATH environment variable.
      2. go env -w GO111MODULE=on --> Go command looks for the go.mod file in the project directory to determine the required dependencies and their versions. If the file exists, the command downloads the required dependencies and stores them in a local cache, which can be shared between projects.
      3. go env -w GO111MODULE=auto --> Enable Go module system if a go.mod file is present in the project directory. If no go.mod file is found, the legacy GOPATH mode is used.
10. Go project structure example:
        project/
        ├── go.mod
        ├── go.sum
        ├── README.md
        ├── cmd/
        │   ├── main.go
        │   └── main_test.go
        ├── internal/
        │   ├── pkg/
        │   │   └── util/
        │   │       └── util.go
        │   └── service/
        │       └── user/
        │           └── user.go
        └── test/
            └── user_test.go