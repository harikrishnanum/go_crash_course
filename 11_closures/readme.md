In Go, a closure is a function value that references variables from outside its body. Specifically, the closure binds variables in the outer scope, so they can be accessed and modified by the function even when it is invoked outside that scope. 

In other words, a closure is a function that has access to variables in its lexical scope, even when the function is called outside of that scope. This allows for more flexible and powerful programming constructs, such as callbacks and iterators.

Here is an example of a closure in Go:

``` go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    a := adder()
    fmt.Println(a(1)) // prints 1
    fmt.Println(a(2)) // prints 3
    fmt.Println(a(3)) // prints 6
}
```

In this example, the `adder` function returns a closure that captures the variable `sum` in its lexical scope. The returned function takes an integer argument, adds it to the `sum` variable, and returns the updated sum. Each time the returned function is called, it updates the `sum` variable, which is preserved between calls because it is part of the closure's captured environment. This allows us to create a stateful function that can be used as an iterator, for example.