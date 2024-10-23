# Go Basic Data Types

This section explains the basic data types available in Go, with examples of how to declare and use them in your programs.

## Table of Contents
- [Introduction](#introduction)
- [Primitive Data Types](#primitive-data-types)
    - [Integer Types](#integer-types)
    - [Unsigned Integer Types](#unsigned-integer-types)
    - [Float Types](#float-types)
    - [Boolean Type](#boolean-type)
    - [String Type](#string-type)
- [Type Conversions](#type-conversions)
- [Examples](#examples)

---

## Introduction

Go has a rich set of basic or **primitive data types** that allow you to work with different types of values, such as numbers, strings, and booleans. Each variable in Go must be declared with a specific type, either explicitly or inferred by the compiler.

---

## Primitive Data Types

### Integer Types

Go provides several types of integers to suit different needs based on size and whether they allow negative values.

- **int**: This is the default integer type, typically 32 or 64 bits depending on the system architecture.
- **int8**: 8-bit signed integer (range: -128 to 127)
- **int16**: 16-bit signed integer (range: -32,768 to 32,767)
- **int32**: 32-bit signed integer (range: -2,147,483,648 to 2,147,483,647)
- **int64**: 64-bit signed integer (range: large range of values)

### Syntax:

```go
var num int = 42
var smallNum int8 = -5
```

### Example:

```go
package main

import "fmt"

func main() {
    var age int = 25
    var small int8 = -100

    fmt.Printf("Age: %d\nSmall: %d\n", age, small)
}
```

---

### Unsigned Integer Types

Unsigned integers only store non-negative values.

- **uint**: Default unsigned integer, usually 32 or 64 bits.
- **uint8**: 8-bit unsigned integer (range: 0 to 255)
- **uint16**: 16-bit unsigned integer (range: 0 to 65,535)
- **uint32**: 32-bit unsigned integer (range: 0 to 4,294,967,295)
- **uint64**: 64-bit unsigned integer (very large range)

### Example:

```go
package main

import "fmt"

func main() {
    var population uint = 1000000
    var smallPopulation uint16 = 500

    fmt.Printf("Population: %d\nSmall Population: %d\n", population, smallPopulation)
}
```

---

### Float Types

Go provides floating-point types to handle decimal numbers.

- **float32**: 32-bit floating-point number
- **float64**: 64-bit floating-point number (default for floating-point values)

### Syntax:

```go
var price float32 = 49.99
var largeNumber float64 = 123456789.123456789
```

### Example:

```go
package main

import "fmt"

func main() {
    var price float32 = 99.99
    var large float64 = 1.2345678912345678

    fmt.Printf("Price: %.2f\nLarge Number: %.8f\n", price, large)
}
```

---

### Boolean Type

Go provides the `bool` type for representing true or false values.

- **bool**: Can hold `true` or `false`.

### Example:

```go
package main

import "fmt"

func main() {
    var isActive bool = true
    var isExpired bool = false

    fmt.Printf("Active: %t\nExpired: %t\n", isActive, isExpired)
}
```

---

### String Type

A **string** is a sequence of characters enclosed within double quotes. Strings in Go are immutable, meaning once a string is created, it cannot be changed.

### Example:

```go
package main

import "fmt"

func main() {
    var greeting string = "Hello, Go!"
    fmt.Println(greeting)
}
```

Strings can also include special characters like newlines (`\n`) and tabs (`\t`).

### Multi-line String Example:

```go
package main

import "fmt"

func main() {
    var multiLine string = `This is a
multi-line
string.`

    fmt.Println(multiLine)
}
```

---

## Type Conversions

Go is a strongly-typed language, which means you cannot automatically assign values of one type to another. Explicit type conversion is required.

### Syntax:

```go
convertedVariable := dataType(variable)
```

### Example:

```go
package main

import "fmt"

func main() {
    var smallNumber int = 5
    var largeNumber float64 = float64(smallNumber) // Convert int to float64

    fmt.Printf("Small: %d, Large: %.2f\n", smallNumber, largeNumber)
}
```

Another important thing to remember is that Go does not automatically convert between types (like an `int` and `float`). You need to manually cast the types to match, as shown above.

---

## Examples

### Example 1: Declaring and Using Different Data Types

```go
package main

import "fmt"

func main() {
    var age int = 30
    var price float64 = 12.50
    var isAvailable bool = true
    var productName string = "Laptop"

    fmt.Printf("Age: %d\nPrice: %.2f\nAvailable: %t\nProduct: %s\n", age, price, isAvailable, productName)
}
```

**Output:**

```plaintext
Age: 30
Price: 12.50
Available: true
Product: Laptop
```

### Example 2: Type Conversion

```go
package main

import "fmt"

func main() {
    var length int = 10
    var width float64 = 7.5

    // Convert 'length' from int to float64 before multiplication
    area := float64(length) * width

    fmt.Printf("Area: %.2f\n", area)
}
```

**Output:**

```plaintext
Area: 75.00
```

---

## Conclusion

Understanding data types is crucial when working with Go. Whether you're dealing with integers, floating-point numbers, booleans, or strings, knowing how to declare and manipulate these types will help you write more efficient and bug-free code. Explicit type conversion ensures that Go maintains its strong typing principles, avoiding implicit type errors.

