Conditional statements in Go are used to control the flow of execution based on certain conditions. Go provides if, if-else, and switch statements for implementing conditional logic.

The if statement in Go works similar to other programming languages, where the code inside the if block is executed only if the given condition is true. The syntax for the if statement is:

``` go
if condition {
    // code to be executed if condition is true
}
```

The if-else statement is used when there are two possible outcomes based on a condition. If the condition is true, the code inside the if block is executed, otherwise, the code inside the else block is executed. The syntax for the if-else statement is:

``` go
if condition {
    // code to be executed if condition is true
} else {
    // code to be executed if condition is false
}
```

Go also provides if-else-if ladder to check multiple conditions in a sequence. The syntax for if-else-if is:

``` go
if condition1 {
    // code to be executed if condition1 is true
} else if condition2 {
    // code to be executed if condition2 is true
} else {
    // code to be executed if all conditions are false
}
```

The switch statement in Go is used to evaluate an expression and compare it against a list of possible cases. If a case matches the value of the expression, the code inside that case block is executed. The syntax for the switch statement is:

``` go
switch expression {
case value1:
    // code to be executed if expression equals value1
case value2:
    // code to be executed if expression equals value2
default:
    // code to be executed if expression doesn't match any case
}
```

In addition to these basic conditional statements, Go also supports some advanced features such as the ternary operator, which can be used to write simple if-else conditions in a more concise manner.