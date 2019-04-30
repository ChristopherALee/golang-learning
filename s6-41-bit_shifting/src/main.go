package main

import "fmt"

func main() {
	x := 2
	// decimal 2xTab binary
	fmt.Printf("%d\t\t%b\n", x, x) // 2 10

	y := x << 1
	// decimal 2xTab binary
	fmt.Printf("%d\t\t%b\n", y, y) // 4 100

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	main2()
}

// use iota to accomplish the above kb/mb/gb
const (
	_   = iota
	kb2 = 1 << (iota * 10)
	mb2 = 1 << (iota * 10)
	gb2 = 1 << (iota * 10)
)

func main2() {
	fmt.Printf("%d\t\t\t%b\n", kb2, kb2)
	fmt.Printf("%d\t\t\t%b\n", mb2, mb2)
	fmt.Printf("%d\t\t%b\n", gb2, gb2)
}
