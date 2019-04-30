package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // false
	x = true
	fmt.Println(x) // true

	a := 7
	b := 42
	fmt.Println(a == b) // false
}
