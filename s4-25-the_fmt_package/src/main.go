package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	// type value
	fmt.Printf("%T\n", y) // int

	// base 2
	fmt.Printf("%b\n", y) // 1000

	// base 16, with lower-case letters for a-f hexidecimal
	fmt.Printf("%x\n", y) // 8

	// hexidecimal alternate format: add leading 0 for octal (%#o), 0x for hex (%#x);
	// 0X for hex (%#X); suppress 0x for %p (%#p);
	// for %q, print a raw (backquoted) string if strconv.CanBackquote
	// returns true;
	fmt.Printf("%#x\n", y) // 0x8

	// multi-value print formatting
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	// String printing
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println(s)

	// print value in default format
	fmt.Printf("%v\n", y)
}
