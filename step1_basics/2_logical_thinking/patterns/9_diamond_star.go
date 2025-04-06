package main

import "fmt"

/*
N = 5

		   *
		  ***
		 *****
		*******
	   *********
	   *********
		*******
		 *****
		  ***
		   *
*/
func main() {
	N := 30
	starCnt := 1
	var flip bool
	for i := 0; i < 2*N; i++ {
		for k := 0; k < ((((N * 2) - 1) - starCnt) / 2); k++ {
			fmt.Printf(" ")
		}

		for k := 0; k < starCnt; k++ {
			fmt.Printf("*")
		}

		for k := 0; k < ((((N * 2) - 1) - starCnt) / 2); k++ {
			fmt.Printf(" ")
		}
		if !flip {
			starCnt += 2
		} else {
			starCnt -= 2
		}

		fmt.Println()

		if starCnt == (2*N)-1 {
			flip = true
		}

	}

}
