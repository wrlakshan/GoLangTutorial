# Go: Loops

This section explains the concept of loops in Go, focusing on the `for` loop, including its syntax, variations, and examples.

## Table of Contents
- [Introduction to Loops](#introduction-to-loops)
- [The `for` Loop](#the-for-loop)
- [Different Forms of `for` Loops](#different-forms-of-for-loops)
- [Using `range` with `for`](#using-range-with-for)
- [Breaking and Continuing Loops](#breaking-and-continuing-loops)
- [Examples](#examples)
- [Conclusion](#conclusion)

---

## Introduction to Loops

Loops are fundamental constructs in programming that allow you to execute a block of code repeatedly based on a condition. In Go, the `for` loop is the only looping construct available and it can be used in various forms.

---

## The `for` Loop

The `for` loop in Go has a simple syntax:

### Syntax:

```go
for initialization; condition; post {
    // Code to be executed
}
```

- **Initialization**: Executes before the loop starts (typically used for variable declaration).
- **Condition**: Evaluated before each iteration. If true, the loop continues; if false, it exits.
- **Post**: Executed at the end of each iteration (commonly used to increment a counter).

### Example:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration:", i) // Output: Iteration: 0, 1, 2, 3, 4
    }
}
```

---

## Different Forms of `for` Loops

### 1. **Traditional `for` Loop**

```go
for i := 0; i < 5; i++ {
    // Code block
}
```

### 2. **Infinite Loop**

You can create an infinite loop by omitting the condition:

```go
for {
    // This will run indefinitely
}
```

### 3. **Using a Condition Only**

You can also omit the initialization and post statements, keeping only the condition:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

---

## Using `range` with `for`

The `range` keyword is used to iterate over elements of various data structures, such as arrays, slices, maps, and strings.

### Syntax:

```go
for index, value := range collection {
    // Code to be executed
}
```

### Example:

```go
package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}

    for index, fruit := range fruits {
        fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
    }
}
```

---

## Breaking and Continuing Loops

You can control the flow of loops using the `break` and `continue` statements.

### `break`

The `break` statement exits the loop immediately.

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit the loop when i is 5
    }
    fmt.Println(i)
}
```

### `continue`

The `continue` statement skips the current iteration and proceeds to the next iteration.

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Output: 1, 3, 5, 7, 9
}
```

---

## Examples

### Example 1: Basic `for` Loop

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello, World!")
    }
}
```

### Example 2: Infinite Loop

```go
package main

import "fmt"

func main() {
    count := 0
    for {
        if count == 5 {
            break // Exit the loop after 5 iterations
        }
        fmt.Println("Count:", count)
        count++
    }
}
```

### Example 3: Using `range` with an Array

```go
package main

import "fmt"

func main() {
    numbers := [3]int{1, 2, 3}

    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### Example 4: Using `continue` in a Loop

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Skip even numbers
        }
        fmt.Println("Odd Number:", i) // Output: Odd Number: 1, 3, 5, 7, 9
    }
}
```

---

## Conclusion

Loops are a fundamental part of programming in Go. The `for` loop, with its various forms and capabilities, provides a powerful way to iterate through data structures and control the flow of execution. Understanding how to effectively use loops is crucial for writing efficient Go programs.

