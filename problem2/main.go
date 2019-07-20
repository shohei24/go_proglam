package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	var y int
	var Ans int
	var operate string

	n, y, Ans = 0, 0, 0

	fmt.Scan(&n)
	fmt.Scan(&operate)
	fmt.Scan(&y)

	fmt.Print(n, operate, y, "\r\n")

	fmt.Printf("Import \r\n")
	switch operate {
	case "-":
		Ans = n - y
	case "+":
		Ans = n + y
	case "*":
		Ans = n * y
	case "/":
		Ans = n / y
	default:
		fmt.Printf("ERROR \r\n")
		os.Exit(1)
	}

	fmt.Print("", Ans)

}
