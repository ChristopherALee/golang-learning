package main

import "fmt"

type sandwich int

var x sandwich

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 8
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}