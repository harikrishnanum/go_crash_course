package main

import "fmt"

// name := "Brad" // This will throw an error as it is outside of a function. Short declaration can only be used inside a function.

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3
	// var unusedVar = "This is an unused variable" // This will throw an error as it is not used anywhere in the program.

	// Shorthand
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)

}
