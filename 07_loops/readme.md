In Go, there are three types of loops: `for`, `while`, and `do-while`.

### For loop
The `for` loop in Go is similar to the `for` loop in other programming languages. It has three components: an initializer, a condition, and a post statement. The initializer is executed only once at the beginning of the loop. The condition is checked before each iteration of the loop. If the condition is true, the body of the loop is executed. The post statement is executed at the end of each iteration.

Here is the syntax for a `for` loop in Go:

```go
for initializer; condition; post {
    // loop body
}
```

### While loop
In Go, the `for` loop can be used as a `while` loop by omitting the initializer and the post statement. 

Here is an example:

```go
for condition {
    // loop body
}
```

### Do-while loop
Go doesn't have a built-in `do-while` loop. However, you can create a `do-while` loop using a `for` loop and a `break` statement.

Here is an example:

```go
for {
    // loop body
    if condition {
        break
    }
}
```

In this example, the loop will continue executing as long as the `condition` is false. Once the `condition` becomes true, the `break` statement is executed, which terminates the loop.