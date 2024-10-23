# Go Programming: Advanced Documentation

This guide provides a comprehensive overview of core concepts and advanced features in Go programming, from basic variable declarations to more advanced topics such as goroutines and channels. Use this as a reference to get started with Go and to understand its most important features.

---

## Table of Contents
- [Module and Package](#module-and-package)
- [Variable Declaration](#variable-declaration)
- [Data Types](#data-types)
- [Pointers](#pointers)
- [User Input](#user-input)
- [Arrays and Slices](#arrays-and-slices)
- [Control Structures](#control-structures)
    - [Loops](#loops)
    - [Conditionals](#conditionals)
- [Functions](#functions)
    - [Receiver Functions](#receiver-functions)
- [Packages in Go](#packages-in-go)
- [Maps](#maps)
- [Structs](#structs)
- [Concurrency](#concurrency)
    - [Goroutines](#goroutines)
    - [Channels](#channels)
- [Scope](#scope)

---

## Module and Package

To create a new module:

```bash
go mod init <module-path>
```

Example:

```bash
go mod init booking-app
```

To run Go files:

```bash
go run main.go
go run main.go bill.go
go run .
```

---

## Variable Declaration

In Go, variables can be declared in multiple ways:

```go
var name string = "test"
const version = "1.0.0"
var age string
city := "Colombo City"
```

You can also print variables using `fmt.Printf`:

```go
fmt.Printf("Name: %v\nVersion: %s\n", name, version)
```

### Basic Data Types
- **String**
- **Integer**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Integer**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Float**: `float32`, `float64`
- **Boolean**: `true`, `false`

---

## Pointers

Pointers allow you to reference the memory address of variables:

```go
&name // Returns the memory address of the variable 'name'
```

---

## User Input

To capture user input:

```go
fmt.Scan(&name)
```

---

## Arrays and Slices

### Arrays

Arrays are fixed-size, index-based collections:

```go
var booking [5]string
booking = [5]string{"foo", "bar"}
```

Access elements by index:

```go
booking[0] // First element
```

### Slices

Slices are dynamic arrays, an abstraction over arrays:

```go
var bookings []string
bookings = append(bookings, "New Booking")
```

---

## Control Structures

### Loops

Go has only one looping construct: the `for` loop.

#### Basic `for` Loop:

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

#### Range-based Loop:

```go
for index, value := range bookings {
    fmt.Printf("Index: %d, Value: %s\n", index, value)
}
```

### Conditionals

#### `if-else` Statement:

```go
if condition {
    // Logic
} else {
    // Alternate Logic
}
```

#### `switch` Statement:

```go
switch city {
    case "a":
        // Logic for case 'a'
    case "b":
        // Logic for case 'b'
    default:
        // Default logic
}
```

---

## Functions

Functions in Go can have multiple return values.

#### Basic Function:

```go
func add(a int, b int) int {
    return a + b
}
```

#### Function with Multiple Return Values:

```go
func getDetails() (string, int) {
    return "Name", 25
}
```

### Receiver Functions

Receiver functions allow you to attach methods to types.

```go
type Bill struct {
    id    int
    items map[string]float64
    name  string
}

func (b *Bill) format() string {
    return fmt.Sprintf("Bill #%d for %s", b.id, b.name)
}
```

---

## Packages in Go

Go packages are a collection of `.go` files that are compiled together. A typical folder structure:

```plaintext
helper/
    - helper.go
main.go
go.mod
```

### Exporting and Importing

- To **export** a function or variable, start its name with an uppercase letter.
- Import packages using the `import` keyword.

---

## Maps

Maps are key-value data structures, but they cannot have mixed data types:

```go
var dates map[string]string
dates = make(map[string]string)

dates["key"] = "value"
```

For a list of maps:

```go
var datesList = make([]map[string]string, 0)
```

---

## Structs

Structs hold mixed data types and define a custom type.

```go
type User struct {
    name string
    age  int
}

var userData = User{
    name: "wrlakshan",
    age:  21,
}
```

---

## Concurrency

Go provides concurrency support with **goroutines** and **channels**.

### Goroutines

A goroutine is a lightweight thread managed by the Go runtime. Use the `go` keyword to run a function concurrently.

```go
go functionName()
```

To ensure the main thread waits for goroutines to finish, use `sync.WaitGroup`:

```go
var wg sync.WaitGroup

func main() {
    wg.Add(1)
    go task()
    wg.Wait() // Ensures main thread waits for goroutines
}

func task() {
    // Task logic
    wg.Done()
}
```

### Channels

Channels are used to share data between goroutines.

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

value := <-ch
fmt.Println(value) // 42
```

---

## Scope

There are three levels of scope in Go:
1. **Local Scope**: Variables declared inside functions.
2. **Package Scope**: Variables declared outside functions are accessible throughout the package.
3. **Global Scope**: Exported variables (capitalized) are accessible across packages.

Example:

```go
var MyVar = "Global Variable" // Global due to capitalization
```

---

## Summary

This guide serves as a starting point for both beginners and advanced Go developers. As you explore further, dive into Go’s concurrency model, efficient memory management, and its use of powerful built-in data structures to build robust applications.

