package main

import "fmt"

// use 'iota' when you need something to auto-increment by 1

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)        // 0
	fmt.Println(b)        // 1
	fmt.Println(c)        // 2
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // int
	fmt.Printf("%T\n", c) // int
	fmt.Println(d)        // 0
	fmt.Println(e)        // 1
	fmt.Println(f)        // 2
}
