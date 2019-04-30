package main

import "fmt"

func main() {
	s := "Hello world"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// convert to slice of bytes
	bs := []byte(s)
	fmt.Println(bs)        // [72 101 108 108 111 32 119 111 114 108 100]
	fmt.Printf("%T\n", bs) // []uint8

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", i) // utf-i codepoint
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}
	// 0 72
	// 1 101
	// 2 108
	// 3 108
	// 4 111
	// 5 32
	// 6 119
	// 7 111
	// 8 114
	// 9 108
	// 10 100
}
