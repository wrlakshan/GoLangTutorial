# Go: Getting User Input

This section explains how to get user input in Go using standard input functions, including reading values from the command line and handling input errors.

## Table of Contents
- [Introduction](#introduction)
- [Using `fmt.Scan`](#using-fmtscan)
- [Using `fmt.Scanln`](#using-fmtscanln)
- [Using `fmt.Scanf`](#using-fmtscanf)
- [Handling Input Errors](#handling-input-errors)
- [Examples](#examples)

---

## Introduction

In Go, you can get user input through the console using functions from the `fmt` package. The most common functions are `fmt.Scan`, `fmt.Scanln`, and `fmt.Scanf`, which allow you to read various types of input from the user, store them in variables, and perform actions based on the input.

---

## Using `fmt.Scan`

The `fmt.Scan` function reads input from standard input (usually the terminal) and stores it in the provided variable(s). It reads input until it encounters a space, newline, or end of input.

### Syntax:

```go
fmt.Scan(&variable)
```

### Example:

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name) // Read user input and store it in the 'name' variable
    fmt.Printf("Hello, %s!\n", name)
}
```

In this example, `fmt.Scan` waits for the user to type their name and hit enter. It stores the input in the `name` variable and then prints it.

---

## Using `fmt.Scanln`

The `fmt.Scanln` function reads user input until it encounters a newline character. This is useful when you want to make sure the input ends with a newline, especially when reading multiple values.

### Syntax:

```go
fmt.Scanln(&variable1, &variable2, ...)
```

### Example:

```go
package main

import "fmt"

func main() {
    var firstName, lastName string
    fmt.Print("Enter your first and last name: ")
    fmt.Scanln(&firstName, &lastName) // Read two values separated by space
    fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}
```

This example reads the first name and last name from the user and stores them in `firstName` and `lastName` respectively. The user must press enter after typing both names.

---

## Using `fmt.Scanf`

The `fmt.Scanf` function allows you to read formatted input, similar to how `fmt.Printf` formats output. You can specify a format string to match the expected input.

### Syntax:

```go
fmt.Scanf(formatString, &variable1, &variable2, ...)
```

### Example:

```go
package main

import "fmt"

func main() {
    var age int
    var city string
    fmt.Print("Enter your age and city (format: age city): ")
    fmt.Scanf("%d %s", &age, &city) // Read age as int and city as string
    fmt.Printf("Age: %d, City: %s\n", age, city)
}
```

In this example, `fmt.Scanf` reads input in the format of an integer followed by a string, separated by a space. The input is stored in the `age` and `city` variables accordingly.

---

## Handling Input Errors

Whenever you use `fmt.Scan`, `fmt.Scanln`, or `fmt.Scanf`, they return the number of successfully scanned items and an error if there’s any issue with the input. You should always check for errors when accepting user input.

### Example:

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    _, err := fmt.Scan(&age) // Check for errors
    if err != nil {
        fmt.Println("Error reading input:", err)
    } else {
        fmt.Printf("Your age is %d\n", age)
    }
}
```

In this example, if the user enters a non-numeric value for age, an error will be printed.

---

## Examples

### Example 1: Reading a Single Value

```go
package main

import "fmt"

func main() {
    var city string
    fmt.Print("Enter your city: ")
    fmt.Scan(&city) // Reading input into the 'city' variable
    fmt.Printf("You live in %s\n", city)
}
```

**Output:**

```plaintext
Enter your city: Colombo
You live in Colombo
```

### Example 2: Reading Multiple Values with `Scanln`

```go
package main

import "fmt"

func main() {
    var firstName, lastName string
    fmt.Print("Enter your first and last name: ")
    fmt.Scanln(&firstName, &lastName) // Reading multiple values
    fmt.Printf("Hello, %s %s!\n", firstName, lastName)
}
```

**Output:**

```plaintext
Enter your first and last name: Alice Johnson
Hello, Alice Johnson!
```

### Example 3: Formatted Input with `Scan`

```go
package main

import "fmt"

func main() {
    var age int
    var city string
    fmt.Print("Enter your age and city (e.g., 25 NewYork): ")
    fmt.Scanf("%d %s", &age, &city) // Reading formatted input
    fmt.Printf("You are %d years old and live in %s.\n", age, city)
}
```

**Output:**

```plaintext
Enter your age and city (e.g., 25 NewYork): 28 London
You are 28 years old and live in London.
```

### Example 4: Handling Input Errors

```go
package main

import "fmt"

func main() {
    var age int
    fmt.Print("Enter your age: ")
    if _, err := fmt.Scan(&age); err != nil {
        fmt.Println("Invalid input. Please enter a number.")
    } else {
        fmt.Printf("Your age is: %d\n", age)
    }
}
```

**Output when an invalid input is given:**

```plaintext
Enter your age: abc
Invalid input. Please enter a number.
```

---

## Conclusion

Getting user input is straightforward in Go using functions from the `fmt` package. Depending on your needs, you can use `Scan`, `Scanln`, or `Scanf` to capture single or multiple values, as well as formatted input. Always remember to handle errors, especially when dealing with user input, to ensure your program remains robust.

