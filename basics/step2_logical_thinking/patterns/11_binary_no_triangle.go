package main

import (
	"fmt"
)

/*
N = 5
1      		i=0, print 1 (start = 1)
0 1   		i=1, print 0 1 (start = 0) (then flip and print till =i)
1 0 1       i=2, print 1 0 1 (start = 1)
0 1 0 1
1 0 1 0 1
*/
func main() {
	n := 6
	inSb := 1

	for i := 0; i < n; i++ {
		sb := inSb
		for j := 0; j <= i; j++ {
			if sb == 1 {
				fmt.Printf("1 ")
				sb = 0
			} else {
				fmt.Printf("0 ")
				sb = 1
			}

		}
		if inSb == 0 {
			inSb = 1
		} else {
			inSb = 0
		}

		fmt.Println()
	}

}
