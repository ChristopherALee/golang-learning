package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s) // H

	// byte
	bs := []byte(s)
	fmt.Println(bs) // [72]

	n := bs[0]
	// decimal
	fmt.Println(n)

	fmt.Printf("%T\n", n)

	// binary
	fmt.Printf("%b\n", n) // 1001000

	// hexidecimal
	fmt.Printf("%#X\n", n) // 48
}
