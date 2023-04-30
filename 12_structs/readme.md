In Go, a structure is a composite data type that groups together zero or more named values of any type. A structure, also known as a struct, is defined using the `type` and `struct` keywords, followed by the name of the structure, and a set of curly braces `{}` that enclose the definition of its fields.

Each field is specified as a name followed by a type, separated by a space. The type can be any valid Go type, including other structures or interfaces. Fields can also have tags, which are used to attach metadata to the field.

Here is an example of defining a structure in Go:

```go
type Person struct {
    name string
    age  int
    address Address
}

type Address struct {
    street string
    city   string
    state  string
}
```

In this example, we have defined a `Person` structure with two fields: `name` of type `string` and `age` of type `int`. The `Person` structure also includes a field `address` of type `Address`, which is itself a structure with three fields.

Once a structure is defined, we can create instances of the structure using the `var` keyword or a shorthand notation:

```go
var person1 Person
person2 := Person{name: "Alice", age: 30, address: Address{street: "123 Main St", city: "New York", state: "NY"}}
```

We can also access the fields of a structure using the dot notation:

```go
fmt.Println(person2.name)
fmt.Println(person2.address.city)
```

Structures can be used to group related data together and pass them as a single unit, making the code more organized and easier to read.

In Go, a method is a function that is associated with a specific type of struct. A method has a receiver, which is a variable of the type the method is associated with. The receiver appears between the func keyword and the name of the method.

The syntax for defining a method is:

``` go
func (receiverName ReceiverType) MethodName(args) {
    // method implementation
}
```

The receiver is a parameter that is defined like any other parameter, but with the receiver type preceding the parameter name. The receiver type is a struct type, and the receiver name is typically a one or two letter abbreviation of the type name.

Here is an example of a struct type with a method:

``` go
type Person struct {
    firstName string
    lastName  string
    age       int
}

func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s %s\n", p.firstName, p.lastName)
}
```

In this example, we define a Person struct with three fields: firstName, lastName, and age. We also define a method SayHello() that takes a receiver of type Person. This method simply prints out a greeting using the values of the firstName and lastName fields.

To call a method, we create a variable of the struct type and use the dot notation to call the method on that variable:

``` go
p := Person{firstName: "John", lastName: "Doe", age: 30}
p.SayHello() // Output: Hello, my name is John Doe
```

In this example, we create a Person variable p and call the SayHello() method on that variable. The SayHello() method uses the values of the firstName and lastName fields of the p variable to print out a greeting.

The receiver argument can be either a value receiver or a pointer receiver. 

A value receiver is a copy of the value passed to the method. This means that if the method modifies the value, the original value outside the method remains unchanged. 

A pointer receiver, on the other hand, receives a pointer to the original value, allowing the method to modify the original value outside the method.

Here is an example of a struct with both value and pointer receiver methods:

``` go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (r *Rectangle) scaleWidth(scale float64) {
    r.width *= scale
}

```
The `area()` method has a value receiver, which means that it receives a copy of the `Rectangle` struct. The `scaleWidth()` method has a pointer receiver, which means that it receives a pointer to the `Rectangle` struct, allowing it to modify the original struct.