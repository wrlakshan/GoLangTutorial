# Go Modules and Packages

This section covers how to initialize a Go module, manage dependencies, and work with packages in Go.

## Table of Contents
- [Introduction to Go Modules](#introduction-to-go-modules)
- [Creating a Module](#creating-a-module)
- [Running Go Programs](#running-go-programs)
- [Understanding Packages](#understanding-packages)
- [Example: Basic Package Structure](#example-basic-package-structure)

---

## Introduction to Go Modules

A **Go module** is a collection of Go packages stored in a directory with a `go.mod` file at its root. Modules are how Go manages dependencies. Every Go project is considered a module, and the `go.mod` file is used to track the project and its dependencies.

Modules simplify dependency management, allowing you to version your code, import third-party libraries, and ensure consistent builds.

---

## Creating a Module

To create a new Go module, you use the `go mod init` command followed by the module path. The module path typically matches the URL where your project is hosted (e.g., GitHub).

### Command:
```bash
go mod init <module-path>
```

### Example:

```bash
go mod init booking-app
```

This will create a `go.mod` file in your project directory, which will look something like this:

```plaintext
module booking-app

go 1.20
```

The `go.mod` file specifies the module name and the version of Go used.

---

## Running Go Programs

Once you have set up the Go module, you can run your Go programs in several ways:

### Running a Single File:

```bash
go run main.go
```

This command runs `main.go` as a standalone program.

### Running Multiple Files:

```bash
go run main.go bill.go
```

This command runs both `main.go` and `bill.go` files together, assuming they are part of the same package.

### Running the Entire Module:

If all your `.go` files are in the same package, you can run them with:

```bash
go run .
```

This command runs all files in the current directory that belong to the package `main`.

---

## Understanding Packages

A **Go package** is a collection of Go files that are compiled together. Every Go file belongs to a package, which is declared at the top of the file:

```go
package main
```

The `main` package is a special package that serves as the entry point of an executable Go program.

### Key Points about Packages:
- **Packages group related code**: Packages make it easy to organize and reuse code across projects.
- **Packages are importable**: Once you create a package, you can import it into other Go files.
- **The `main` package**: This package is the starting point of a Go program. Any Go file with the `main` package and a `main()` function can be run directly.

---

## Example: Basic Package Structure

### Folder Structure:

```plaintext
booking-app/
├── go.mod
├── main.go
└── helper/
    └── helper.go
```

### `main.go`:

```go
package main

import (
	"fmt"
	"booking-app/helper" // Importing custom package
)

func main() {
	fmt.Println("Welcome to the Booking App")
	helper.DisplayHelp() // Calling a function from another package
}
```

### `helper/helper.go`:

```go
package helper

import "fmt"

func DisplayHelp() {
	fmt.Println("This is the help section for booking.")
}
```

### Running the Program:

To run the entire program, simply use:

```bash
go run .
```

Output:
```plaintext
Welcome to the Booking App
This is the help section for booking.
```

---

## Conclusion

Modules and packages are fundamental to organizing your Go code and managing dependencies. By following the steps outlined in this guide, you can set up your own Go modules and start working with packages to keep your code modular and maintainable.

