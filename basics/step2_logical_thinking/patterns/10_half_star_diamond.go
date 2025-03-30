package main

import "fmt"

/*
N=5

* i=1, 1 time.
** i = 2, 2
*** i =3, 3
**** i=4, 4
***** i=5,5
****  i=6,4 , 6-N,N-1,
***   i=7,3 , 7-N,N-2
**
*
*/
func main() {
	N := 6
	var flip bool
	for i := 1; i < (N * 2); i++ {
		if !flip {
			for j := 0; j < i; j++ {
				fmt.Printf("*")
			}
		}

		if i == N {
			flip = true
			fmt.Println()
			continue
		}
		if flip {
			for j := 2*N - i; j > 0; j-- {
				fmt.Printf("*")

			}
		}

		fmt.Println()

	}

}
