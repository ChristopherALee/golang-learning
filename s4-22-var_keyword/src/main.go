package main

import "fmt"

// var declaration operator used within global scope
var y = 2

// DECLARE there is a variable with the identifier 'z' is of type int
// Assigns the 0 value of type int to 'z'
var z int

func main() {
	// short declaration operator used within function scope
	x := 1
	fmt.Println(x) // 1
	fmt.Println(y) // 2
	foo()          // printed 2 from foo
	fmt.Println(z) // 0
}

func foo() {
	fmt.Println("printed", y, "from foo")
}
