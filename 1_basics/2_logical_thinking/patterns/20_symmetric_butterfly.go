package main

import "fmt"

/*
N = 5
*        *               , i=0, j<=i for *. sp: (2*n) - 2*(i+1) => 2n -2i -2. 2(n-i-1) -> spaces.
**      **
***    ***
****  ****
**********
****  ****               ,i=5, j<n-(i-n), sp: (2*n) - 2*(n-(i-n)) => 2n - 2(2n-i) =>2i-2n
***    ***
**      **
*        *
*/
func main() {
	n := 6
	for i := 1; i < 2*n; i++ {

		if i < n {
			for j := 0; j < i; j++ {
				fmt.Printf("*")
			}
			for j := 0; j < 2*(n-i); j++ {
				fmt.Printf(" ")
			}
			for j := 0; j < i; j++ {
				fmt.Printf("*")
			}
		} else {
			for j := 0; j < 2*n-i; j++ {
				fmt.Printf("*")
			}
			for j := 0; j < 2*(i-n); j++ {
				fmt.Printf(" ")
			}
			for j := 0; j < 2*n-i; j++ {
				fmt.Printf("*")
			}
		}
		fmt.Println()

	}

}
