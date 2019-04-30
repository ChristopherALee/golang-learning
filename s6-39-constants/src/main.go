package main

import "fmt"

const a = 42
const b = 42.78
const c = "Batman"

// OR

const (
	a = 42
	b = 42.78
	c = "Batman"
)

// can also strict type
const (
	a int     = 42
	b float64 = 42.78
	c string  = "Batman"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // float64
	fmt.Printf("%T\n", c) // string
}
