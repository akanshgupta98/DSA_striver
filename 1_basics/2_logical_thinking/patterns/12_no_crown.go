package main

import (
	"fmt"
)

/*
N=5
1      1    , total sp: 2*n-2. sp: (2*n-2)-2*i => 2 ((n-1)-i)
12    21    ,i=2, sp:
123  321
12344321
*/
func main() {

	n := 6
	for i := 1; i < n; i++ {

		// forward loop
		for j := 1; j <= i; j++ {
			fmt.Printf("%v", j)
		}
		// spaces
		for j := 0; j < (2 * ((n - 1) - i)); j++ {
			fmt.Printf(" ")
		}

		// backward loop
		for j := i; j >= 1; j-- {
			fmt.Printf("%v", j)
		}
		fmt.Println()

	}

}
