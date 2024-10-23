# Go: Mutable and Immutable Data Types

This section covers mutable and immutable data types in Go, explaining the differences, examples, and practical applications.

## Table of Contents
- [Introduction to Mutable and Immutable Data Types](#introduction-to-mutable-and-immutable-data-types)
- [Mutable Data Types](#mutable-data-types)
- [Immutable Data Types](#immutable-data-types)
- [Examples of Mutable and Immutable Data Types](#examples-of-mutable-and-immutable-data-types)
- [Conclusion](#conclusion)

---

## Introduction to Mutable and Immutable Data Types

In programming, **mutable** data types can be changed after their creation, while **immutable** data types cannot be modified once they are created. Understanding the distinction between these two categories is crucial for managing state and performance in Go.

---

## Mutable Data Types

In Go, mutable data types include:
- **Slices**
- **Maps**
- **Channels**
- **Structs**

You can change the content of these types without creating a new instance.

### Example: Mutable Slices

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3}
    fmt.Println("Original Slice:", numbers) // Output: [1 2 3]

    numbers[0] = 10 // Modify the first element
    fmt.Println("Modified Slice:", numbers) // Output: [10 2 3]
}
```

### Example: Mutable Maps

```go
package main

import "fmt"

func main() {
    colors := make(map[string]string)
    colors["red"] = "#FF0000" // Add an entry
    fmt.Println("Colors Map:", colors) // Output: map[red:#FF0000]

    colors["red"] = "#FF1111" // Modify the existing entry
    fmt.Println("Updated Colors Map:", colors) // Output: map[red:#FF1111]
}
```

---

## Immutable Data Types

In Go, immutable data types include:
- **Arrays**
- **Strings**

Once these types are created, you cannot change their contents without creating a new instance.

### Example: Immutable Strings

```go
package main

import "fmt"

func main() {
    str := "Hello, World!"
    fmt.Println("Original String:", str) // Output: Hello, World!

    // Attempt to modify the string
    str[0] = 'h' // This will cause a compilation error
}
```

To modify a string, you need to create a new one:

```go
newStr := "h" + str[1:] // Create a new string
fmt.Println("Modified String:", newStr) // Output: hello, World!
```

### Example: Immutable Arrays

```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println("Original Array:", arr) // Output: [1 2 3]

    // Attempt to modify the array
    arr[0] = 10 // This will change the original array
    fmt.Println("Modified Array:", arr) // Output: [10 2 3]
}
```

While you can modify the contents of an array, if you assign it to a new array, it will create a copy:

```go
arr2 := arr // Copy the array
arr2[0] = 100
fmt.Println("Original Array After Copy:", arr) // Output: [10 2 3]
fmt.Println("Copied Array:", arr2) // Output: [100 2 3]
```

---

## Examples of Mutable and Immutable Data Types

### Example 1: Using a Mutable Slice

```go
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    slice = append(slice, 4) // Append an element
    fmt.Println("Slice after append:", slice) // Output: [1 2 3 4]
}
```

### Example 2: Using an Immutable String

```go
package main

import "fmt"

func main() {
    str := "Immutable"
    str = str + " String" // Create a new string
    fmt.Println("New String:", str) // Output: Immutable String
}
```

### Example 3: Using a Mutable Map

```go
package main

import "fmt"

func main() {
    person := make(map[string]int)
    person["Alice"] = 30
    person["Bob"] = 25

    fmt.Println("Person Map:", person) // Output: map[Alice:30 Bob:25]

    person["Alice"] = 31 // Update value
    fmt.Println("Updated Person Map:", person) // Output: map[Alice:31 Bob:25]
}
```

### Example 4: Using an Immutable Array

```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println("Array:", arr) // Output: [1 2 3]

    // Arrays are immutable in terms of size
    // arr = append(arr, 4) // This will cause an error
}
```

---

## Conclusion

Understanding mutable and immutable data types is essential for effective programming in Go. Mutable data types allow for dynamic modifications, while immutable types provide safety and predictability. By leveraging these concepts, you can write more efficient and reliable code.

