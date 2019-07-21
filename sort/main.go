package main

import (
	"fmt"
)

func main() {
	var a [10]int
	a[0] = 1
	a[1] = 10
	a[2] = 11
	a[3] = 21
	a[4] = 9
	a[5] = 5
	a[6] = 131
	a[7] = 15
	a[8] = 31
	a[9] = 22

	//b := [4]int{1, 2, 3, 5}

	var i int
	var k int

	/*bubble sort work up to down*/
	/*Go LANGEGE can't do do-while.
	  and , while Statement equate for Statement
	*/
	for i = len(a) - 2; i >= 0; i-- {
		for k = 0; k <= i; k++ {
			if a[k] > a[k+1] {
				a[k], a[k+1] = a[k+1], a[k]

			}
		}

	}

	for i = 0; i < 10; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Printf("\r\n")

}

/* sumple code
func bubble_sort(ar []int) []int {
    for i := len(ar) - 2 ; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            if ar[j] > ar[j + 1] {
               ar[j], ar[j + 1] = ar[j + 1], ar[j]
            }
        }
    }
    return ar
}
*/
