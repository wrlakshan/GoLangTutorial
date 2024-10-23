# Go: Functions

This section covers the concept of functions in Go, including their definition, parameters, return values, and examples of different types of functions.

## Table of Contents
- [Introduction to Functions](#introduction-to-functions)
- [Defining Functions](#defining-functions)
- [Function Parameters](#function-parameters)
- [Return Values](#return-values)
- [Multiple Return Values](#multiple-return-values)
- [Anonymous Functions](#anonymous-functions)
- [Defer Statement](#defer-statement)
- [Examples](#examples)
- [Conclusion](#conclusion)

---

## Introduction to Functions

Functions are a fundamental building block in Go programming. They allow you to encapsulate code, making it reusable, organized, and easier to read. Functions can accept parameters and return values, enabling you to perform operations on data.

---

## Defining Functions

A function in Go is defined using the `func` keyword, followed by the function name, parameters, and the return type (if any).

### Syntax:

```go
func functionName(parameterType parameterName) returnType {
    // Code to execute
}
```

### Example:

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello, " + name + "!") // Output: Hello, John!
}

func main() {
    greet("John")
}
```

---

## Function Parameters

You can define parameters for functions, allowing you to pass data into them. Parameters are specified in parentheses after the function name.

### Example:

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("Sum:", result) // Output: Sum: 8
}
```

You can also omit the type for the second parameter if both parameters are of the same type:

```go
func add(a, b int) int {
    return a + b
}
```

---

## Return Values

Functions can return a single value or multiple values. The return type(s) are specified after the parameter list.

### Example:

```go
package main

import "fmt"

func subtract(a int, b int) int {
    return a - b
}

func main() {
    result := subtract(10, 4)
    fmt.Println("Difference:", result) // Output: Difference: 6
}
```

---

## Multiple Return Values

Go allows functions to return multiple values, which is useful for returning both results and error states.

### Example:

```go
package main

import "fmt"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Quotient:", result) // Output: Quotient: 5
    }
}
```

---

## Anonymous Functions

Anonymous functions are functions defined without a name. They can be used for callbacks, closures, or encapsulating logic.

### Example:

```go
package main

import "fmt"

func main() {
    add := func(a, b int) int {
        return a + b
    }

    result := add(2, 3)
    fmt.Println("Sum:", result) // Output: Sum: 5
}
```

You can also use anonymous functions immediately:

```go
package main

import "fmt"

func main() {
    func() {
        fmt.Println("Hello from an anonymous function!") // Output: Hello from an anonymous function!
    }()
}
```

---

## Defer Statement

The `defer` statement postpones the execution of a function until the surrounding function returns. It is often used for cleanup activities, like closing files or releasing resources.

### Example:

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Goodbye!")
    fmt.Println("Hello!") // Output: Hello!
    // Goodbye! will be printed when main returns
}
```

You can defer multiple functions; they will execute in last-in-first-out order:

```go
package main

import "fmt"

func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    fmt.Println("Hello!") // Output: Hello!
    // Output order of defer: Second, First
}
```

---

## Examples

### Example 1: Basic Function

```go
package main

import "fmt"

func multiply(x int, y int) int {
    return x * y
}

func main() {
    result := multiply(3, 4)
    fmt.Println("Product:", result) // Output: Product: 12
}
```

### Example 2: Function with Multiple Return Values

```go
package main

import "fmt"

func getDetails() (string, int) {
    return "Alice", 30
}

func main() {
    name, age := getDetails()
    fmt.Printf("Name: %s, Age: %d\n", name, age) // Output: Name: Alice, Age: 30
}
```

### Example 3: Using Anonymous Function

```go
package main

import "fmt"

func main() {
    square := func(n int) int {
        return n * n
    }

    result := square(5)
    fmt.Println("Square:", result) // Output: Square: 25
}
```

### Example 4: Using `defer`

```go
package main

import "fmt"

func main() {
    defer fmt.Println("This will run at the end.")
    fmt.Println("This will run first.") // Output: This will run first.
}
```

---

## Conclusion

Functions are essential in Go for structuring code, making it modular, and allowing for code reuse. Understanding how to define functions, use parameters, handle return values, and apply concepts like anonymous functions and deferred execution is crucial for writing effective Go programs.

