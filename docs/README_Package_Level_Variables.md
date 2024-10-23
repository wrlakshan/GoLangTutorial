# Go: Package-Level Variables

This section covers package-level variables in Go, detailing their declaration, scope, and examples of usage.

## Table of Contents
- [Introduction to Package-Level Variables](#introduction-to-package-level-variables)
- [Declaring Package-Level Variables](#declaring-package-level-variables)
- [Scope of Package-Level Variables](#scope-of-package-level-variables)
- [Examples](#examples)
- [Conclusion](#conclusion)

---

## Introduction to Package-Level Variables

Package-level variables are variables declared outside of any function in a Go package. They are accessible to all functions within the same package, making them useful for sharing state or configuration across multiple functions.

---

## Declaring Package-Level Variables

To declare a package-level variable, simply define it outside any function, typically at the top of your file. You can declare it with or without an initial value.

### Syntax:

```go
var variableName variableType
var variableName = initialValue
```

### Example:

```go
package main

import "fmt"

// Package-level variable
var appName string = "My Application"

func main() {
    fmt.Println("App Name:", appName) // Output: App Name: My Application
}
```

---

## Scope of Package-Level Variables

Package-level variables have a scope that extends throughout the entire package in which they are declared. Any function within that package can access and modify these variables.

### Example:

```go
package main

import "fmt"

// Package-level variable
var counter int = 0

// Function to increment the counter
func increment() {
    counter++
}

// Function to print the counter
func printCounter() {
    fmt.Println("Counter:", counter)
}

func main() {
    increment()
    printCounter() // Output: Counter: 1
}
```

In the example above, the `counter` variable is accessed and modified by both the `increment` and `printCounter` functions.

---

## Examples

### Example 1: Basic Package-Level Variable

```go
package main

import "fmt"

// Package-level variable
var message string = "Hello, World!"

func main() {
    fmt.Println(message) // Output: Hello, World!
}
```

### Example 2: Modifying Package-Level Variable

```go
package main

import "fmt"

// Package-level variable
var score int = 0

// Function to add points
func addPoints(points int) {
    score += points
}

func main() {
    addPoints(10)
    fmt.Println("Score:", score) // Output: Score: 10
    addPoints(5)
    fmt.Println("Score:", score) // Output: Score: 15
}
```

### Example 3: Using Package-Level Variables in Multiple Functions

```go
package main

import "fmt"

// Package-level variable
var userName string

// Function to set user name
func setUser(name string) {
    userName = name
}

// Function to greet the user
func greetUser() {
    fmt.Println("Hello,", userName)
}

func main() {
    setUser("Alice")
    greetUser() // Output: Hello, Alice
}
```

---

## Conclusion

Package-level variables in Go are a powerful feature that allows for sharing state across multiple functions within a package. They are particularly useful for maintaining configuration, counters, or shared data that needs to be accessed by various parts of your application.
