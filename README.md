# goTutorial
## useful links
1. https://go.dev/
2. https://pkg.go.dev/std --> Standard library
3. https://go-proverbs.github.io/ --> Do & Don't
4. testing package --> https://pkg.go.dev/testing@go1.21.3#T
## go commands
1. go mod --> module maintenance, will display several commands; go help mod (The get all the options)
   1. go mod init <NAME OF MODULE>--> initialize a new module in current directory, create a go.mod file (it's a configuration file about the module and it dependencies) inside the module name is demo.
      1. Example: go mod init ps.m3.demo1 
      2. Example: go mod init demo
      3. Example: go mod init github.com/pluralsight/gomodules
   2. go mod tidy --> remove unused dependencies from go.mod file.
   3. go mod vendor -v --> Create a vendor folder in the same folder(root folder), it will contain the packages that the app use.
   4. go mod verfiy --> verify that when we build the module all the dependencies modules are with the right version.
   5. go mod why <package> --> Why do we need a specifc package that exists in our go.mod file
   6. go mod graph --> display the dependencies graph for the application
   7. go mod download <value> --> download a package/module into the go cache, it will not display in the go.mod file, but it will be available for us if we are without internet and we use the go get command, Example: go mod download github.com/pioz/faker@master
   8. go mod edit <parameter> <value > --> Edit the go.mod file, can be used with a tool 
      1. Example: go mod edit -module <value> --> set the **module** paramater with the <value>
      2. Example: go mod edit -go <value> --> set the **go** paramater with the <value>
      3. Example: go mod edit -require <value> --> add or update the **require** paramater with the <value>
      4. Example: go mod edit -droprequire <value> --> drop the <value> from the **require** paramater
      5. Example: go mod edit -exclude <value> --> add the package/module <value> in the **exclude** parameter.
      6. Example: go mod edit -dropexclude <value> --> drop the package/module <value> from the **exclude** parameter.
      7. Example: go mod edit -replace <value>=../quote --> **replace** the package/module <value> with the code from the quote directory(we cloned the package/module code into the quote folder)
      8. Example: go mod edit -dropreplace <value> --> drop the **replace** parameter
      9. Example: go mod edit -print --> prints the content of the go.mod file
      10. Example: go mod edit -json --> prints the content of the go.mod file in a json format.
2. go build --> compile the code and create a binary file in the directory.
   1. Example: go build . --> In (unix) <moduleName> file will appear. (windows) <moduleName>.exe apper
   2. Example: go build -mod=vendor . --> build the app with the dependencies from the vendor folder.
   3. Example: go build -mod=readonly . --> build the app it will fail if the same version will exists in the require and exlude parameter. If we not use the readonly the go tool will change the version in the reqire parameter automatically
3. go run . --> compile the current code into a temp directory and run the Go program.
   1. go run <moduleName> --> Run the binary compiled go file
   2. go run -mod=vendor . --> Run the app but use the dependencies from the vendor folder
4. go get <package> --> Thie will download a package
   1. Example: go get golang.org/x/exp/slices
   2. Example: go get github.com/pioz/faker
   3. Example: go get rsc.io/quote -->(https://github.com/rsc/quote) Go will default the version 1, although there could be version 2. We can check the versions with the go list command 6.3 section
   4. Example: go get rsc.io/quote/v2 --> Get the latest release of version v2 
   5. Example: go get github.com/pioz/faker@v1.6.0 - download a specific version
   6. Example: go get github.com/pioz/faker@<COMMIT_NUM> - download a specific commit
   7. Example: go get github.com/pioz/faker@latest - download a latest tag, it is also the default when we don't defined aversion we will get the latest release (or Tag) version of the package 
   8. Example: go get github.com/pioz/faker@master - download a latest commit of the master branch
   9. Example: go get github.com/pioz/faker@>=1.7.2 - download a version equal or greater then 1.7.2
   10. Example go get -u --> Automatically updated the packages to latest version
5. go test <Tests locations> 
   1.  go test . --> Run all tests in root directory
   2.  go test ./knowledge --> Run all tests in knowledge directory
   3.  got test ./...  --> Run all tests in root directory and subsirectories (All test in the module)
6. go list --> display the go module name that we are working on.
   1. go list all --> display all the packages that exists in our project.
   2. go list -m all --> We get only the modules that we use, in our app
   3. go list -m -versions <ModuleName> --> Display all avilable version for module.
      1. Example go list -m -versions golang.org/x/exp
7.  go env --> display all go variables, 
   1. Example GOPATH variable where packages are located
   2. Example GOMODCACHE where the packages are stored in the cache
   3. go parameter GO111MODULE --> Used when we want to use am external tool to manage modules dependencies
      1. go env -w GO111MODULE=off --> Remove the Go module awareness means go command looks for packages in the directories specified by the GOPATH environment variable.
      2. go env -w GO111MODULE=on --> Go command looks for the go.mod file in the project directory to determine the required dependencies and their versions. If the file exists, the command downloads the required dependencies and stores them in a local cache, which can be shared between projects.
      3. go env -w GO111MODULE=auto --> Enable Go module system if a go.mod file is present in the project directory. If no go.mod file is found, the legacy GOPATH mode is used.
## Go project structure
1.  Go project structure example:
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
## Go versioning
1.  Semantic Versioning v1.5.3-pre1
   1.  v - version prefix (required)
   2.  1 - Major revision (likely to break backword compatability)
   3.  5 - Minor revision (new features, doesn't break BC)
   4.  3 - Patch (bug fixes, no new features, and doesn't breadk BC)
   5.  pre1 - Pre-release of new version, if applicable (text is arbitrary)