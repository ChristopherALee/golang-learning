package main

import "fmt"

func main() {
	g := 1 == 1
	h := 1 <= 2
	i := 1 >= 0
	j := 1 != 0
	l := 2 > 1

	a := [5]bool{g, h, i, j, l}

	for i, s := range a {
		fmt.Println(i, s)
	}
}
