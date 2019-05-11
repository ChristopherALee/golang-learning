package main

import "fmt"

func main() {
	a := 1
	// decimal, binary, hex
	fmt.Printf("%d\t%b\t%#x", a, a, a)

	fmt.Println()

	// shift by 1
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
