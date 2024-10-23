# Go: Maps

This section covers maps in Go, detailing their structure, how to create and manipulate them, and examples of usage.

## Table of Contents
- [Introduction to Maps](#introduction-to-maps)
- [Declaring and Initializing Maps](#declaring-and-initializing-maps)
- [Accessing and Modifying Maps](#accessing-and-modifying-maps)
- [Iterating Over Maps](#iterating-over-maps)
- [Maps with Different Data Types](#maps-with-different-data-types)
- [Example](#example)
- [Conclusion](#conclusion)

---

## Introduction to Maps

In Go, a map is a built-in data type that represents a collection of key-value pairs. Maps are unordered and allow for fast lookups, additions, and deletions. They are similar to dictionaries or hash tables in other programming languages.

---

## Declaring and Initializing Maps

To declare a map, you can use the built-in `make` function or a map literal. The key and value types must be specified.

### Syntax:

1. Using `make`:

```go
myMap := make(map[KeyType]ValueType)
```

2. Using a map literal:

```go
myMap := map[KeyType]ValueType{
    key1: value1,
    key2: value2,
}
```

### Example:

```go
package main

import "fmt"

func main() {
    // Using make
    ages := make(map[string]int)

    // Using map literal
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }

    fmt.Println(ages)   // Output: map[]
    fmt.Println(colors) // Output: map[blue:#0000FF green:#00FF00 red:#FF0000]
}
```

---

## Accessing and Modifying Maps

You can access and modify map values using their keys. If the key does not exist, a new entry will be created with the default value for the value type.

### Accessing a Value:

```go
value := myMap[key]
```

### Modifying a Value:

```go
myMap[key] = newValue
```

### Deleting a Key-Value Pair:

Use the `delete` function:

```go
delete(myMap, key)
```

### Example:

```go
package main

import "fmt"

func main() {
    // Initialize a map
    studentGrades := make(map[string]string)

    // Add key-value pairs
    studentGrades["Alice"] = "A"
    studentGrades["Bob"] = "B"

    // Access a value
    fmt.Println("Alice's Grade:", studentGrades["Alice"]) // Output: Alice's Grade: A

    // Modify a value
    studentGrades["Bob"] = "A+"
    fmt.Println("Bob's Updated Grade:", studentGrades["Bob"]) // Output: Bob's Updated Grade: A+

    // Delete a key-value pair
    delete(studentGrades, "Alice")
    fmt.Println("Student Grades after deletion:", studentGrades) // Output: Student Grades after deletion: map[Bob:A+]
}
```

---

## Iterating Over Maps

You can use a `for` loop with the `range` keyword to iterate over the keys and values in a map.

### Syntax:

```go
for key, value := range myMap {
    // Do something with key and value
}
```

### Example:

```go
package main

import "fmt"

func main() {
    // Initialize a map
    capitals := map[string]string{
        "USA":   "Washington, D.C.",
        "France": "Paris",
        "Japan": "Tokyo",
    }

    // Iterate over the map
    for country, capital := range capitals {
        fmt.Printf("The capital of %s is %s\n", country, capital)
    }
}
```

Output:

```
The capital of USA is Washington, D.C.
The capital of France is Paris
The capital of Japan is Tokyo
```

---

## Maps with Different Data Types

Maps can store any data type as values, and the keys can be of any comparable type (e.g., strings, integers).

### Example:

```go
package main

import "fmt"

func main() {
    // Map with integer keys and string values
    intMap := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
    }

    fmt.Println(intMap) // Output: map[1:One 2:Two 3:Three]
}
```

---

## Example

Here’s a complete example demonstrating various operations on maps:

```go
package main

import "fmt"

func main() {
    // Create a map
    phoneBook := map[string]string{
        "John":  "123-456-7890",
        "Jane":  "098-765-4321",
        "Alice": "555-555-5555",
    }

    // Accessing a value
    fmt.Println("John's Phone Number:", phoneBook["John"]) // Output: John's Phone Number: 123-456-7890

    // Modifying a value
    phoneBook["John"] = "111-111-1111"
    fmt.Println("Updated John's Phone Number:", phoneBook["John"]) // Output: Updated John's Phone Number: 111-111-1111

    // Adding a new entry
    phoneBook["Bob"] = "222-222-2222"
    fmt.Println("Phone Book:", phoneBook)

    // Deleting an entry
    delete(phoneBook, "Alice")
    fmt.Println("Phone Book after deletion:", phoneBook)

    // Iterating over the map
    for name, number := range phoneBook {
        fmt.Printf("%s's Phone Number: %s\n", name, number)
    }
}
```

---

## Conclusion

Maps in Go provide a powerful way to store and manage collections of key-value pairs. They are versatile and allow for quick lookups, updates, and deletions. Understanding how to declare, manipulate, and iterate over maps is essential for effective Go programming.

