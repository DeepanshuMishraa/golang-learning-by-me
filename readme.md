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
