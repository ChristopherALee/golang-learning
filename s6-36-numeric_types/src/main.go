package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

func main() {
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on.
	fmt.Println(runtime.GOOS)

	// GOOS is the running program's operating system target: one of darwin, freebsd, linux, and so on. To view possible combinations of GOOS and GOARCH, run "go tool dist list".
	fmt.Println(runtime.GOARCH)
}
