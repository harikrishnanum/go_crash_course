In Go, a goroutine is a lightweight thread of execution managed by the Go runtime. Goroutines are designed to be easy to use and efficient, allowing developers to write concurrent code with ease.

To create a new goroutine, you can use the `go` keyword followed by a function call:

```go
go myFunc()
```

This will create a new goroutine that executes the `myFunc` function concurrently with the main program. The main program and the new goroutine will execute independently of each other, allowing them to perform different tasks simultaneously.

Goroutines are very lightweight, so it's possible to create many of them in a single program without using too much memory or CPU resources. They are also very efficient, as they can be multiplexed onto a smaller number of OS threads, allowing the Go runtime to manage them more efficiently.

Here's an example of how to use goroutines in Go to perform some work concurrently:

```go
func main() {
    go func() {
        fmt.Println("This is running in a goroutine")
    }()

    fmt.Println("This is running in the main function")

    // Wait for a key press before exiting
    fmt.Scanln()
}
```

In this example, we create a new goroutine that prints a message to the console. We then print another message from the main function. Because the goroutine is running concurrently, we will see the two messages appear in the console at the same time.

Goroutines are a powerful feature of Go that enable developers to write highly concurrent and efficient programs. They can be used to perform background tasks, handle network connections, and much more. However, they do require careful management, as they can lead to complex synchronization issues if not used correctly.

## Concurrency vs parallelism
Concurrency and parallelism are related but distinct concepts in computer science. They both involve executing multiple tasks simultaneously, but they differ in how they achieve this.

Concurrency refers to the ability of a program to perform multiple tasks at the same time, without necessarily running them simultaneously. This is typically achieved by interleaving the execution of different tasks, such that they appear to be running simultaneously. This interleaving can be managed through the use of threads, processes, or coroutines.

Parallelism, on the other hand, refers to the ability of a program to perform multiple tasks simultaneously, by actually running them on different processors or cores in a multi-core system. Parallelism is achieved by dividing a large task into smaller sub-tasks that can be executed simultaneously, and then combining the results once all sub-tasks have completed.

In other words, concurrency is a property of a program's design, while parallelism is a property of the underlying hardware. A program can be concurrent without being parallel, but a program that is parallel is also concurrent.

In Go, concurrency is achieved through the use of goroutines and channels, which allow developers to write highly concurrent programs with ease. Parallelism can be achieved in Go by using the `runtime` package to control the number of operating system threads used by the Go scheduler.

To summarize, concurrency is about managing the interleaving of tasks, while parallelism is about executing tasks on multiple processors or cores simultaneously.

<hr>

The main.go is a simple Go program that listens for TCP connections on port 9000 of the localhost and responds to each connection by sending the current time every second until the connection is closed.

Here's a breakdown of how the program works:

1. First, the program imports several standard Go packages including `fmt`, `io`, `net`, and `os`.

2. The `main` function starts by calling the `net.Listen` function, which creates a new TCP listener that listens for incoming connections on port 9000 of the localhost. If an error occurs during the creation of the listener, the program prints the error to the console and exits.

3. The program then enters an infinite loop that listens for new connections using the `listener.Accept` function. For each new connection, a new goroutine is spawned to handle the connection.

4. The new goroutine writes the current time to the connection every second using the `io.WriteString` function. If an error occurs during the write, the goroutine returns, which causes the connection to be closed.

5. The program continues to listen for new connections and spawn new goroutines to handle them until the program is terminated.

### Running the program
To run the program, you can use the `go run` command:

```bash
go run main.go
```
Then open a new terminal and use the `nc` command to connect to the server:

```bash
nc localhost 9000
```