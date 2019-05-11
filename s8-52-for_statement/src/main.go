package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	y := 1
	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y++
	}
}
