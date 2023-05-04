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
In Go, the `switch` statement does not require an explicit `break` statement after each case. In Go, once a matching case is executed, the control automatically exits the switch block, unless there is a `fallthrough` statement explicitly used, in which case control is transferred to the next case.

`fallthrough` is a keyword in Go that is used to control the flow of control in a switch statement. When `fallthrough` is used, control is transferred to the next case block, even if it does not match the condition, i.e., it "falls through" the cases.

In Go, each case statement in a switch statement is automatically terminated by a `break` statement, which means that when the code in a case block is executed, the switch statement will exit immediately. However, if the `fallthrough` keyword is used at the end of a case block, control will be transferred to the next case block, regardless of whether its condition is true or false.

Here is an example:

``` go
switch x {
case 1:
	fmt.Println("One")
	fallthrough
case 2:
	fmt.Println("Two")
case 3:
	fmt.Println("Three")
default:
	fmt.Println("Unknown")
}
```

In the above code, if the value of `x` is 1, the output will be:
```
One
Two
```
because the `fallthrough` keyword causes the execution to fall through to the next case block, which is case 2. If `fallthrough` was not used, the output would be only "One".

In addition to these basic conditional statements, Go also supports some advanced features such as the ternary operator, which can be used to write simple if-else conditions in a more concise manner.
