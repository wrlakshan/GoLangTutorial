# Go: Packages

This section covers packages in Go, detailing their structure, how to create and import them, and examples of usage.

## Table of Contents
- [Introduction to Packages](#introduction-to-packages)
- [Creating a Package](#creating-a-package)
- [Importing Packages](#importing-packages)
- [Exporting Functions and Variables](#exporting-functions-and-variables)
- [Example Package Structure](#example-package-structure)
- [Using Third-Party Packages](#using-third-party-packages)
- [Conclusion](#conclusion)

---

## Introduction to Packages

In Go, a package is a collection of Go source files that are compiled together. Packages are the fundamental way to organize and reuse code in Go applications. Each Go program is made up of packages, with the `main` package serving as the entry point.

---

## Creating a Package

To create a package, you need to create a directory that contains one or more Go source files. Each source file should start with a `package` declaration, followed by the package name.

### Example:

1. Create a new directory called `mathutils`.
2. Inside `mathutils`, create a file named `mathutils.go`.

#### `mathutils.go`:

```go
package mathutils

// Add function to sum two integers
func Add(a int, b int) int {
    return a + b
}
```

---

## Importing Packages

To use a package in your Go application, you need to import it using the `import` statement. The import path is typically the relative or absolute path to the package directory.

### Example:

In your `main` package, you can import the `mathutils` package:

#### `main.go`:

```go
package main

import (
    "fmt"
    "path/to/your/mathutils" // Adjust the import path as necessary
)

func main() {
    result := mathutils.Add(3, 4)
    fmt.Println("Result:", result) // Output: Result: 7
}
```

Make sure to adjust the import path according to your project structure.

---

## Exporting Functions and Variables

In Go, only functions and variables that start with an uppercase letter are exported from a package, meaning they can be accessed from other packages.

### Example:

In the `mathutils` package:

```go
package mathutils

// Exported function
func Add(a int, b int) int {
    return a + b
}

// Unexported function
func subtract(a int, b int) int {
    return a - b
}
```

In the `main` package, you can call `Add`, but not `subtract`:

```go
package main

import (
    "fmt"
    "path/to/your/mathutils"
)

func main() {
    sum := mathutils.Add(5, 2) // Accessible
    // diff := mathutils.subtract(5, 2) // Not accessible, would cause a compile error
    fmt.Println("Sum:", sum) // Output: Sum: 7
}
```

---

## Example Package Structure

Here’s an example folder structure for a Go project using packages:

```
my-go-project/
│
├── main.go               // Main application file
│
└── mathutils/            // Custom package
    └── mathutils.go      // Source file for the mathutils package
```

### `main.go`:

```go
package main

import (
    "fmt"
    "my-go-project/mathutils"
)

func main() {
    fmt.Println("Sum:", mathutils.Add(10, 5)) // Output: Sum: 15
}
```

---

## Using Third-Party Packages

You can use third-party packages in your Go projects by downloading them using the `go get` command. For example:

```bash
go get github.com/some/package
```

After downloading, you can import it into your Go files:

```go
package main

import (
    "fmt"
    "github.com/some/package"
)

func main() {
    // Use functions from the third-party package
}
```

---

## Conclusion

Packages in Go provide a modular way to organize and reuse code. Understanding how to create, import, and export functions and variables in packages is essential for building maintainable Go applications. Additionally, leveraging third-party packages can significantly enhance your project's capabilities.
