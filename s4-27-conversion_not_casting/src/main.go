package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)        // 42
	fmt.Printf("%T\n", a) // int
	fmt.Println(b)        // 43
	fmt.Printf("%T\n", b) // main.hotdog

	// CANNOT do this. Two different types.
	// STATIC programming language
	// 'a' is already declared of type INT
	// 'b' is declared of type HOTDOG
	// a = b

	// HOWEVER
	// take the value of type HOTDOG and convert to type INT
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
