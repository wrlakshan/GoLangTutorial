# Go: Receiver Functions

This section covers receiver functions in Go, detailing how they allow you to define methods on types and how to use value and pointer receivers.

## Table of Contents
- [Introduction to Receiver Functions](#introduction-to-receiver-functions)
- [Defining Receiver Functions](#defining-receiver-functions)
- [Value Receivers vs. Pointer Receivers](#value-receivers-vs-pointer-receivers)
- [Examples](#examples)
- [Conclusion](#conclusion)

---

## Introduction to Receiver Functions

In Go, receiver functions allow you to associate methods with types, effectively enabling object-oriented programming paradigms. A receiver can be a value or a pointer of a specific type, allowing the method to operate on instances of that type.

---

## Defining Receiver Functions

A receiver function is defined by specifying a receiver argument in the function signature. The receiver can be either a value or a pointer.

### Syntax:

```go
func (receiverType receiverName) methodName(parameters) returnType {
    // Code to execute
}
```

### Example:

```go
package main

import "fmt"

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func main() {
    circle := Circle{radius: 5}
    fmt.Println("Area of circle:", circle.area()) // Output: Area of circle: 78.5
}
```

---

## Value Receivers vs. Pointer Receivers

### Value Receivers

When a method is defined with a value receiver, a copy of the receiver is made. This is useful when you do not need to modify the original instance or when the receiver type is small (like a basic data type).

### Example:

```go
package main

import "fmt"

type Rectangle struct {
    width, height float64
}

// Method with value receiver
func (r Rectangle) area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 4, height: 5}
    fmt.Println("Area of rectangle:", rect.area()) // Output: Area of rectangle: 20
}
```

### Pointer Receivers

When a method is defined with a pointer receiver, the method can modify the original instance. This is preferred for larger structs or when you want to change the state of the receiver.

### Example:

```go
package main

import "fmt"

type Counter struct {
    count int
}

// Method with pointer receiver
func (c *Counter) increment() {
    c.count++
}

func main() {
    counter := Counter{count: 0}
    counter.increment()
    fmt.Println("Counter:", counter.count) // Output: Counter: 1
}
```

---

## Examples

### Example 1: Method with Value Receiver

```go
package main

import "fmt"

type Book struct {
    title  string
    author string
}

// Method with value receiver
func (b Book) details() string {
    return b.title + " by " + b.author
}

func main() {
    book := Book{title: "1984", author: "George Orwell"}
    fmt.Println(book.details()) // Output: 1984 by George Orwell
}
```

### Example 2: Method with Pointer Receiver

```go
package main

import "fmt"

type Person struct {
    name string
}

// Method with pointer receiver
func (p *Person) changeName(newName string) {
    p.name = newName
}

func main() {
    person := Person{name: "Alice"}
    person.changeName("Bob")
    fmt.Println("Updated Name:", person.name) // Output: Updated Name: Bob
}
```

### Example 3: Using Both Value and Pointer Receivers

```go
package main

import "fmt"

type Square struct {
    side float64
}

// Method with value receiver
func (s Square) area() float64 {
    return s.side * s.side
}

// Method with pointer receiver
func (s *Square) setSide(newSide float64) {
    s.side = newSide
}

func main() {
    square := Square{side: 4}
    fmt.Println("Area of square:", square.area()) // Output: Area of square: 16

    square.setSide(5)
    fmt.Println("Updated Area of square:", square.area()) // Output: Updated Area of square: 25
}
```

---

## Conclusion

Receiver functions in Go provide a powerful way to define methods on custom types, allowing for organized and modular code. Understanding the difference between value and pointer receivers is essential for effective memory management and method behavior in Go applications.
