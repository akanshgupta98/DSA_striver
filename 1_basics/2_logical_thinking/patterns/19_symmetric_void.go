package main

import "fmt"

/*
N=5
**********    i=0,j=5-0 to 0, sp: n*2 - 2*(n-i) = 2n - 2n + 2i
****  ****    i=1,j=5-1 to 0,sp
***    ***
**      **
*        *
*        *  i=5,j=5-5 to 0, sp: n*2 - 2(i-n) = 2n - 2i+2n = 4n - 2i -> 2(2n-i) = sp: 2 * (10-5) = 10.
**      **
***    ***
****  ****
**********
*/
func main() {
	n := 10
	// var flip bool

	for i := 0; i <= 2*n; i++ {
		// Second half
		if i >= n {
			// print star.
			for j := i - n; j >= 0; j-- {
				fmt.Printf("*")

			}
			// print space.
			for j := 0; j < 2*(2*n-i); j++ {
				fmt.Printf(" ")
			}
			// print star.
			for j := i - n; j >= 0; j-- {
				fmt.Printf("*")
			}

		} else { // First Half
			// print star.
			for j := n - i; j >= 0; j-- {
				fmt.Printf("*")
			}

			// print space.
			for j := 0; j < 2*i; j++ {
				fmt.Printf(" ")
			}
			// print star.
			for j := n - i; j >= 0; j-- {
				fmt.Printf("*")
			}
		}

		fmt.Println()

	}

}
