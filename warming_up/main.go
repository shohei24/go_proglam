package main

import (
	"fmt"
)

func main() {
	var n int
	var y int
	n, y = 0, 3

	for n < 10 { // Qestion 1
		n++
		fmt.Print(" ", n)

	}
	fmt.Print("\r\n") // Questio 2
	fmt.Scan(&y)
	fmt.Print(" y =  ", y)
	y = y * 3
	fmt.Print("\r\n y =  ", y)

}
