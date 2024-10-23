# Go Variable Declaration

This section provides an overview of how to declare variables in Go, covering different types of variable declarations, constants, and examples.

## Table of Contents
- [Introduction](#introduction)
- [Basic Variable Declaration](#basic-variable-declaration)
- [Short Variable Declaration](#short-variable-declaration)
- [Constants](#constants)
- [Zero Values](#zero-values)
- [Type Inference](#type-inference)
- [Printing Variables](#printing-variables)
- [Examples](#examples)

---

## Introduction

In Go, variables can be declared in multiple ways. This flexibility makes it easier to manage variable scopes and types based on the situation.

---

## Basic Variable Declaration

The most common way to declare variables in Go is using the `var` keyword followed by the variable name and its type.

### Syntax:

```go
var variableName variableType
```

### Example:

```go
var name string
var age int
```

Here, `name` is a string variable, and `age` is an integer variable. If no value is assigned at the time of declaration, the variable will have a default **zero value**.

You can also initialize a variable at the time of declaration:

```go
var city string = "Colombo"
```

Alternatively, Go can infer the type based on the initial value:

```go
var country = "Sri Lanka"  // Go infers the type as string
```

---

## Short Variable Declaration

In Go, there’s a shorthand way to declare and initialize variables using the `:=` operator. This is commonly used within functions.

### Syntax:

```go
variableName := value
```

### Example:

```go
name := "Alice"
age := 30
```

The type of the variable is automatically inferred based on the value on the right-hand side.

---

## Constants

Constants are declared using the `const` keyword. A constant's value cannot be changed once assigned.

### Syntax:

```go
const constantName = value
```

### Example:

```go
const version = "1.0.0"
const pi = 3.14159
```

Unlike variables, constants cannot be declared using the `:=` shorthand notation, and they must be explicitly assigned when declared.

---

## Zero Values

In Go, uninitialized variables are automatically given a **zero value** depending on their type:

- **int**: `0`
- **string**: `""` (empty string)
- **bool**: `false`
- **pointers**: `nil`

### Example:

```go
var age int        // age = 0
var active bool    // active = false
var name string    // name = ""
```

---

## Type Inference

Go can infer the type of a variable from its initial value, so you don't always need to specify the type. This is particularly useful with the `:=` syntax.

### Example:

```go
name := "Bob"  // Go infers the type as string
age := 25      // Go infers the type as int
```

---

## Printing Variables

To print variables in Go, you use `fmt.Printf` or `fmt.Println`.

### Example of `fmt.Printf`:

```go
fmt.Printf("Name: %s\nAge: %d\n", name, age)
```

The `%s` is a placeholder for strings, and `%d` is for integers. You can use different format specifiers based on the type of the variable.

### Example of `fmt.Println`:

```go
fmt.Println("Hello, " + name)
```

---

## Examples

### Example 1: Basic Variable Declaration and Initialization

```go
package main

import "fmt"

func main() {
    var name string = "John"
    var age int = 25
    var city = "New York" // Type inferred as string

    fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", name, age, city)
}
```

**Output:**

```plaintext
Name: John
Age: 25
City: New York
```

### Example 2: Using Short Declaration

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    city := "Los Angeles"

    fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", name, age, city)
}
```

**Output:**

```plaintext
Name: Alice
Age: 30
City: Los Angeles
```

### Example 3: Declaring Constants

```go
package main

import "fmt"

const version = "1.0.0"
const pi = 3.14159

func main() {
    fmt.Printf("Version: %s\nPi: %f\n", version, pi)
}
```

**Output:**

```plaintext
Version: 1.0.0
Pi: 3.141590
```

### Example 4: Zero Values

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    var isMember bool

    fmt.Printf("Name: %q\nAge: %d\nIs Member: %t\n", name, age, isMember)
}
```

**Output:**

```plaintext
Name: ""
Age: 0
Is Member: false
```

---

## Conclusion

Variable declaration is fundamental in Go and provides multiple ways to declare and initialize variables. By understanding the different approaches—whether you are using `var`, `:=`, or `const`—you can manage data more effectively in your programs. Type inference and the concept of zero values make Go both efficient and easy to work with.

