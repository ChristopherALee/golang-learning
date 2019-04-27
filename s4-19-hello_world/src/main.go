package main

import "fmt"

func main() {
	fmt.Println("hello world")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			foo(i)
		} else {
			bar(i)
		}
	}
}

func foo(i int) {
	fmt.Println("foo at", i)
}

func bar(i int) {
	fmt.Println("bar at", i)
}
