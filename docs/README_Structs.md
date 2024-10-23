# Go: Structs

This section covers structs in Go, detailing their structure, how to define and use them, and examples of usage.

## Table of Contents
- [Introduction to Structs](#introduction-to-structs)
- [Defining a Struct](#defining-a-struct)
- [Creating Instances of Structs](#creating-instances-of-structs)
- [Accessing Struct Fields](#accessing-struct-fields)
- [Structs with Methods](#structs-with-methods)
- [Embedding Structs](#embedding-structs)
- [Example](#example)
- [Conclusion](#conclusion)

---

## Introduction to Structs

In Go, a struct is a composite data type that groups together variables (fields) under a single name. Structs are used to create complex data types and are similar to classes in object-oriented programming, but they do not support inheritance.

---

## Defining a Struct

You define a struct using the `type` keyword followed by the struct name and the `struct` keyword. Inside the braces, you specify the fields and their types.

### Syntax:

```go
type StructName struct {
    FieldName1 FieldType1
    FieldName2 FieldType2
}
```

### Example:

```go
type User struct {
    Name string
    Age  int
}
```

---

## Creating Instances of Structs

You can create an instance of a struct using a composite literal. You can initialize fields by name or by position.

### Syntax:

1. By name:

```go
user := User{Name: "Alice", Age: 30}
```

2. By position:

```go
user := User{"Alice", 30}
```

### Example:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of User
    user := User{Name: "Alice", Age: 30}

    // Print user details
    fmt.Println("User:", user)
}
```

---

## Accessing Struct Fields

You can access the fields of a struct using the dot (`.`) operator.

### Example:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

func main() {
    user := User{Name: "Bob", Age: 25}

    // Accessing fields
    fmt.Println("Name:", user.Name) // Output: Name: Bob
    fmt.Println("Age:", user.Age)   // Output: Age: 25
}
```

---

## Structs with Methods

You can define methods on structs. A method is a function that has a receiver, which is a value or pointer of the struct type.

### Syntax:

```go
func (receiver StructType) MethodName() {
    // Method logic
}
```

### Example:

```go
package main

import "fmt"

type User struct {
    Name string
    Age  int
}

// Method to display user information
func (u User) DisplayInfo() {
    fmt.Printf("Name: %s, Age: %d\n", u.Name, u.Age)
}

func main() {
    user := User{Name: "Charlie", Age: 28}
    user.DisplayInfo() // Output: Name: Charlie, Age: 28
}
```

---

## Embedding Structs

Struct embedding allows one struct to include another struct. This provides a way to compose structs and share fields and methods.

### Example:

```go
package main

import "fmt"

type Address struct {
    City    string
    Country string
}

type User struct {
    Name    string
    Age     int
    Address // Embedding Address struct
}

func main() {
    user := User{
        Name:    "David",
        Age:     40,
        Address: Address{City: "New York", Country: "USA"},
    }

    fmt.Println("User:", user.Name)
    fmt.Println("City:", user.City) // Accessing embedded field
    fmt.Println("Country:", user.Country) // Accessing embedded field
}
```

---

## Example

Here’s a complete example demonstrating the use of structs, methods, and embedding:

```go
package main

import "fmt"

// Define structs
type Address struct {
    City    string
    Country string
}

type User struct {
    Name    string
    Age     int
    Address // Embedding Address struct
}

// Method to display user information
func (u User) DisplayInfo() {
    fmt.Printf("Name: %s, Age: %d, City: %s, Country: %s\n", u.Name, u.Age, u.City, u.Country)
}

func main() {
    // Create an instance of User
    user := User{
        Name:    "Eve",
        Age:     35,
        Address: Address{City: "Los Angeles", Country: "USA"},
    }

    // Display user info
    user.DisplayInfo() // Output: Name: Eve, Age: 35, City: Los Angeles, Country: USA
}
```

---

## Conclusion

Structs in Go provide a powerful way to define complex data types and organize related data. Understanding how to define, create, and manipulate structs, as well as how to define methods for them, is essential for effective Go programming. Struct embedding further enhances the capability to compose complex types.

