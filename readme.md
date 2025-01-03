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

### 3. [Variables](variables)

- Variables are used to store values.

- Variables are declared using the `var` keyword followed by the variable name and the type.

- The type of the variable is optional. If the type is not provided, the type is inferred from the value.

```go
var c, python, java bool;
```

The above statement declares three variables c, python, and java of type bool. The default value of a bool variable is false.

Another way to declare variables:

```go
name := "John";
```

This can be used only inside a function. The type of the variable is inferred from the value.Only use this if the initial value is provided and the type is not needed.

```go
var i = 10;
```

Here, the type of the variable is inferred from the value.


### 4. [Data Types](data-types)

- Go is a statically typed language. The type of a variable is known at compile time.

- Go has the following basic data types:

    - bool
    - string
    - int, int8, int16, int32, int64
    - uint, uint8, uint16, uint32, uint64, uintptr
    - byte (alias for uint8)
    - rune (alias for int32)
    - float32, float64
    - complex64, complex128

- The zero value of a variable is the default value of the variable.

- The zero value of a bool variable is false.

- The zero value of a numeric variable is 0.

- The zero value of a string variable is an empty string.

- The zero value of a slice variable is nil.


```go
var(
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)
```

This is another way to declare variables. Here, the type of the variables is explicitly mentioned.

Variable Identifier:

- %T -> type of the variable
- %v -> value of the variable
- %s -> string representation of the value
- %d -> decimal representation of the value

```go
fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
fmt.Printf("Type: %T Value: %v\n", z, z)
```

uint64 -> unsigned 64-bit integers (0 to 18446744073709551615) -> 0 to 2^64 - 1

byte -> alias for uint8
rune -> alias for int32


### printf vs println vs print

- fmt.Printf -> formatted output -> when you want to format the output like C printf

- fmt.Println -> adds a newline at the end of the output

- fmt.Print -> does not add a newline at the end of the output
