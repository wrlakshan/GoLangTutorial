# Go: Goroutines

This section covers goroutines in Go, detailing their purpose, how to use them, and examples of concurrent programming.

## Table of Contents
- [Introduction to Goroutines](#introduction-to-goroutines)
- [Creating Goroutines](#creating-goroutines)
- [Goroutines and the Go Scheduler](#goroutines-and-the-go-scheduler)
- [Synchronization with WaitGroups](#synchronization-with-waitgroups)
- [Channels](#channels)
- [Example](#example)
- [Conclusion](#conclusion)

---

## Introduction to Goroutines

Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions and make it easy to write concurrent programs in Go. Goroutines can run in parallel, leveraging multi-core processors effectively.

---

## Creating Goroutines

To create a goroutine, you use the `go` keyword followed by a function call. The function will run concurrently with the calling function.

### Syntax:

```go
go functionName()
```

### Example:

```go
package main

import (
    "fmt"
    "time"
)

func greet() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go greet() // Create a goroutine

    // Wait for the goroutine to finish
    time.Sleep(time.Second)
    fmt.Println("Main function finished.")
}
```

In this example, the `greet` function runs concurrently with the `main` function. The `Sleep` function ensures that the program waits for the goroutine to finish before exiting.

---

## Goroutines and the Go Scheduler

The Go runtime includes a scheduler that manages the execution of goroutines. It schedules goroutines onto OS threads, allowing efficient use of system resources.

- **Concurrency**: Multiple goroutines can run concurrently, making the program more responsive.
- **Preemption**: The Go scheduler can preempt a running goroutine to allow others to run, ensuring fair execution.

### Example:

```go
package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(time.Millisecond * 500) // Simulate work
    }
}

func main() {
    go printNumbers() // Start goroutine

    // Main function continues to execute
    fmt.Println("Main function is doing other work...")

    // Wait for the goroutine to finish
    time.Sleep(time.Second * 3)
    fmt.Println("Main function finished.")
}
```

---

## Synchronization with WaitGroups

When using goroutines, you may need to wait for them to finish before proceeding. The `sync.WaitGroup` is used for this purpose.

### Example:

```go
package main

import (
    "fmt"
    "sync"
)

func worker(wg *sync.WaitGroup, id int) {
    defer wg.Done() // Notify that the goroutine is done
    fmt.Printf("Worker %d is working...\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Add to the WaitGroup
        go worker(&wg, i) // Start goroutine
    }

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All workers finished.")
}
```

In this example, the `worker` function simulates some work, and the `WaitGroup` ensures that the main function waits until all workers complete.

---

## Channels

Channels are a way to communicate between goroutines. They allow you to send and receive values, enabling synchronization between concurrent processes.

### Creating Channels:

You create a channel using the `make` function.

### Syntax:

```go
ch := make(chan Type)
```

### Sending and Receiving:

- Send: `ch <- value`
- Receive: `value := <-ch`

### Example:

```go
package main

import "fmt"

func ping(ch chan string) {
    ch <- "Ping" // Send value to channel
}

func main() {
    ch := make(chan string) // Create a channel

    go ping(ch) // Start goroutine

    message := <-ch // Receive value from channel
    fmt.Println(message) // Output: Ping
}
```

---

## Example

Here’s a complete example that demonstrates goroutines, synchronization, and channels:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, ch chan string, wg *sync.WaitGroup) {
    defer wg.Done() // Notify that the goroutine is done
    time.Sleep(time.Second) // Simulate work
    msg := fmt.Sprintf("Worker %d finished", id)
    ch <- msg // Send message to channel
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan string)

    for i := 1; i <= 5; i++ {
        wg.Add(1) // Add to the WaitGroup
        go worker(i, ch, &wg) // Start goroutine
    }

    go func() {
        wg.Wait() // Wait for all goroutines to finish
        close(ch) // Close the channel
    }()

    // Receive messages from the channel
    for msg := range ch {
        fmt.Println(msg)
    }

    fmt.Println("All workers finished.")
}
```

In this example, multiple workers are created as goroutines. They send messages back to the main function through a channel. The channel is closed once all workers are done.

---

## Conclusion

Goroutines are a powerful feature in Go that enable concurrent programming. Understanding how to create and manage goroutines, synchronize them with `WaitGroups`, and communicate using channels is essential for building efficient and responsive applications.
