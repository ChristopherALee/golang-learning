package main

import "fmt"

var y = 88

// DECLARED the VARIABLE with the identifier 'z'
// is of TYPE string
// CANNOT REDECLARE of a different type
var z string = "I am a string"

// raw string literal
var a = `z says, "I am a string"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
