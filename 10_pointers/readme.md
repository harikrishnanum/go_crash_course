In Go, a pointer is a variable that holds the memory address of another variable. It allows us to indirectly access and modify the value of a variable.

To declare a pointer variable, we use the asterisk (*) symbol before the type of the variable it points to. For example, to declare a pointer to an integer variable:

``` go
var ptr *int
```

We can also use the `&` operator to get the memory address of a variable. For example:

``` go
var x int = 42
ptr := &x // ptr now holds the memory address of x
```

To access the value of a variable through a pointer, we use the `*` operator. For example:

``` go
var x int = 42
ptr := &x
fmt.Println(*ptr) // prints 42
```

We can also modify the value of a variable through a pointer. For example:

``` go
var x int = 42
ptr := &x
*ptr = 24
fmt.Println(x) // prints 24
```

Pointers are useful in situations where we want to avoid copying large amounts of data. By passing a pointer to a function instead of the data itself, we can modify the original data directly without creating a new copy.

Go supports pass by reference using pointers. When we pass a pointer to a variable, we are passing a reference to the memory location where the variable is stored. Any changes made to the value at that memory location will be reflected in the original variable.

For example:

``` go
func increment(x *int) {
    *x = *x + 1
}

func main() {
    a := 10
    fmt.Println("Before increment: ", a) // prints "Before increment: 10"
    increment(&a)
    fmt.Println("After increment: ", a) // prints "After increment: 11"
}
```

In the above code, we define a function `increment` that takes a pointer to an integer as its argument. The function then increments the value at the memory location pointed to by the pointer. We call the `increment` function by passing it the address of the variable `a` using the `&` operator. The `increment` function modifies the value of `a`, which is then reflected in the output of the `fmt.Println` statement in the `main` function.