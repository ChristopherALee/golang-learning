package main

import "fmt"

func main() {
	y := 0
	for {
		y++

		if y > 100 {
			break
		}

		if y%2 != 0 {
			continue
		}

		fmt.Println(y)
	}

	fmt.Println("done")
}
