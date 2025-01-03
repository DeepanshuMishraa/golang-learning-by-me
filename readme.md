## My Golang Notes

### 1. [Hello World](hello-world)

### Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### What is go.mod file?

- go.mod file is a module definition file. It defines the module path and the dependencies of the module.

- the name of the module should be prefix with hosting site. For example, if you are hosting your module on github, the module name should be github.com/username/module-name.

- go.mod file is created when you run `go mod init module-name` command.

- go.mod file contains the module name and the dependencies of the module.

- we import packages and run the executable file using the module name.

- `go mod download` command is used to download the dependencies of the module that is written in the go.mod file .

- `go mod tidy` command is used to remove the unused dependencies from the go.mod file.

## UTF8 Encoding Support -> https://pkg.go.dev/unicode/utf8

- UTF8 is a variable-length encoding for Unicode code points.

- Golang provides a package called `utf8` to work with UTF8 encoding.

- The `utf8` package provides functions to encode and decode UTF8 encoded strings.

### Running a Go Program

- `go run file-name.go` command is used to run a go program.

- `go build file-name.go` command is used to compile a go program.It builds an executable file.

- `./file-name` command is used to run the compiled go program.

### 2. [Functions](functions)

### Functions -> https://golang.org/doc/effective_go.html#functions

- Functions are the building blocks of a Go program.

- Functions are defined using the `func` keyword followed by the function name, parameters, and return type.

- The return type is optional.

- The function body is enclosed in curly braces `{}`.

An example of a function:

```go
func add(x int, y int) int{
	return x+y;
}
```

Another way to define the function:

```go
func swap(x,y string)(string,string){
	return y,x;
}
```

Here , since we needed to return two strings, we have enclosed the return types in parentheses.


Another way to define the function:

```go
func split(sum int)(x,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
```

Here, we have named the return values x and y. These named return values are treated as variables defined at the top of the function.Not recommended for long functions.


