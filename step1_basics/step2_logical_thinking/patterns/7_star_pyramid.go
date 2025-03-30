package main

import "fmt"

/*
N = 5
    *
   ***
  *****
 *******
*********

======================
i = 0, print star 1 time.
i = 1, print start 3 times.
i = 2, print star 5 times.

col : 9,k=1, 8/2 -> 4. print 4 spaces, 1 star, 4 spaces.
col : 9,k=3,6/2 -> 3, ,print 3 spaces, 3 star, 3 spaces.
col: 9,k=5,4/2 -> 2, print 2 spaces,5 start, 2 spaces.
*/

func main() {
	N := 6
	col := (N * 2) - 1
	k := 1
	for i := 0; i < N; i++ {
		psp := 0
		ps := 0
		for j := 0; j < col; j++ {
			if psp < (col-k)/2 {
				fmt.Printf(" ")
				psp++
				continue
			}
			fmt.Printf("*")
			ps++
			if ps == k {
				psp = 0
			}

		}
		k += 2
		fmt.Println()
	}
}
