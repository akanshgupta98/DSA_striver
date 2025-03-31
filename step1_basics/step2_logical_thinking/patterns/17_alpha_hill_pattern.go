package main

import "fmt"

/*
N = 4

												: 1, 0,6/2
		A          , i=0, total: n*2-1, forward: A to A+i, backward:A+i-1 to A., space: ()
	   ABA         ,i=1, forward: 2,1, (n*2-1) - (i+1+i) / 2
	  ABCBA
	 ABCDCBA
*/
func main() {

	n := 6
	total := 2*n - 1
	for i := 0; i < n; i++ {
		// space.
		for j := 0; j <= ((total - (2*i + 1)) / 2); j++ {
			fmt.Printf(" ")
		}
		// forward.

		for j := 65; j <= 65+i; j++ {
			fmt.Printf("%v", string(j))
		}

		//backward
		for j := 65 + i - 1; j >= 65; j-- {
			fmt.Printf("%v", string(j))
		}
		//space
		for j := 0; j <= ((total - (2*i + 1)) / 2); j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}
