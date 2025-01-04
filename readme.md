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
- %f -> float representation of the value
- %g -> float representation of the value
- %p -> pointer representation of the value
- %q => a single-quoted character literal safely escaped with Go syntax.

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

### Explicit Type Conversion

- Go requires explicit type conversion when converting between different types.

- Type conversion is done using the type name followed by the value in parentheses.

```go
	//type conversion
	var x,y int = 3,4;
	var f float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf("Type %T Value : %v\n",f,f);
	var m uint = uint(f);
	fmt.Printf("Type %T Value : %v\n",m,m);
```

### 5. [Constants](constants)

- Constants are declared using the `const` keyword followed by the constant name and the value.

- Constants cannot be declared using the `:=` syntax.

- Constants can be character, string, boolean, or numeric values.

```go
const Pi = 3.14
```

### 6.Loops

- Go has only one looping construct, the for loop.

- The for loop has three components: the init statement, the condition, and the post statement.

- The init statement is executed before the first iteration.

- The condition is evaluated before every iteration.

- The post statement is executed at the end of every iteration.

- The init statement and the post statement are optional.

- The for loop can be used to iterate over an array, slice, string, map, or channel.

```go
func main(){
    sum := 0;
    for i:=0;i<10;i++{
        sum += i;
    }
    fmt.Println(sum);
}
```

here the init statement is `i:=0`, the condition is `i<10`, and the post statement is `i++`.

#### Implementing a while loop in Go

- Go does not have a while loop.

- The for loop can be used to implement a while loop.

```go
func main(){
    sum := 1;
    for sum < 1000{
        sum += sum;
    }
    fmt.Println(sum);
}
```

Since, there is no init or post statement, it is equivalent to a while loop.

### if else else if statement

- The if statement is used to execute a block of code if the condition is true.

- The else statement is used to execute a block of code if the condition is false.

```go
func main(){
    fmt.Println(sqrt(2),sqrt(-4));
}

func sqrt(x float64) string{
    if x<0{
        return sqrt(-x) + "i";
    }
    return fmt.Sprint(math.Sqrt(x));
}
```

- The if statement can have an optional init statement that is executed before the condition.

- The scope of the variables declared in the init statement is limited to the if block.

```go

func pow(x,n,lim float64) float64{
    if v:= math.Pow(x,n); v<lim{
        return v;
    }
    return lim;
}
```

### Switch Statement

- The switch statement is used to execute a block of code based on the value of an expression.

- The switch statement can be used to replace long if-else chains.

```go
switch os:= runtime.GOOS;os{
	case "darwin":
		fmt.Println("OS X");
	case "linux":
		fmt.Println("Linux");
	default:
		//freebsd,openbsd,windows
		fmt.Printf("%s.\n",os);
	}
```

Here runtime is a package that provides functions to interact with the Go runtime environment.

### Defer Statement

- The defer statement is used to defer the execution of a function until the surrounding function returns.

- The deferred function is executed in LIFO order.

- The deferred function is executed even if the surrounding function panics.

```go
func main(){
    defer fmt.Println("world");
    fmt.Println("hello");
}
```

Here, the `world` statement is deferred until the `main` function returns.

In simple terms, defer is used to delay the execution of a function until the surrounding function returns.

```go
func main(){
	fmt.Println("Counting");

	for i:=0;i<10;i++{
		defer fmt.Println(i);
	}

	fmt.Println("DOne");
}
```

Here the deferred function is executed in LIFO order.
means the output will be 9,8,7,6,5,4,3,2,1,0. as the deferred function is executed after the surrounding function returns. So first the loop will run and then the deferred function will be executed

#### why to use defer?

- Defer is used to ensure that a function call is performed later in a program's execution, usually for purposes of cleanup.

- Defer is commonly used to close a file, unlock a mutex, or print a finalizing message.

- Let's say you open a file and in the end you close the file but the file is not closed due to some error in the middle of the program. So, defer is used to ensure that the file is closed in the end.and avoids the memory leaks. In the end when you defer the closing it will even close the file if there is an error in the middle of the program.

### Reference Types

- Reference types are types that hold references to the underlying data.

- Reference types are used to store the memory address of the value.

- Reference types are used to share data between two or more variables.

- The golang reference types are:

  - slices
  - maps
  - channels
  - pointers

#### Pointers

```go
func main(){
	i,j := 42,2701;

	p := &i;
	fmt.Println(*p); // point to i
	*p = 21; // set i through the pointer->dereferencing
	fmt.Println(i); // see the new value of i

	p = &j; // point to j
	*p = *p / 37; // divide j through the pointer
	fmt.Println(j); // see the new value of j
}
```

Explanation:

- `p := &i` -> p points to the memory address of i.

- `*p = 21` -> set the value of i through the pointer.

- `p = &j` -> p points to the memory address of j.

- `*p = *p / 37` -> divide the value of j through the pointer.

- `fmt.Println(j)` -> see the new value of j.

#### Structs

- A struct is a collection of fields.

- A struct is a user-defined type.

- A struct is used to group related data.

- A struct is defined using the `type` and `struct` keywords.

```go
func main(){

// create a struct

type Vertex struct{
	x int
	y int
}
	//initialise a struct

	v:= Vertex{1,2}
	v.x = 4;
	fmt.Println(v.x)
	fmt.Println(v);

}
```

Here we have created a struct `Vertex` with two fields `x` and `y`. We have initialized the struct with the values `1` and `2`. We have changed the value of `x` to `4`.


- Pointers to structs

- Struct fields can be accessed using a struct pointer.

- To access the field `x` of a struct `v`, we can use `v.x`.

```go
func main(){
    v := Vertex{1,2}
    p := &v
    p.x = 1e9
    fmt.Println(v)
}
```
