package main

import "fmt"

/*
N=4
4 4 4 4 4 4 4     , (0,0):N, (0,1) : N ... (0,2n-1) : N
4 3 3 3 3 3 4     , (1,0):N, (1,1) : N-1,...(1,5):N-1,(1,2n-1):N
4 3 2 2 2 3	4     , (2,0):N, (2,1) : N-1,
4 3	2 1 2 3 4     , ......
4 3	2 2 2 3 4     ,
4 3	3 3 3 3 4     ,
4 4 4 4 4 4 4     , (2n-1,0):N
*/
func main() {
	n := 5
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {

			if i == 0 || j == 0 || i == 2*n-2 || j == 2*n-2 {
				fmt.Printf("%v ", n)
			} else if i == 1 || i == 2*n-2-i || j == 1 || j == 2*n-2-i {
				fmt.Printf("%v ", n-i)
				// } else if i == 2 || i == 2*n-4 || j == 2 || j == 2*n-4 {
				// 	fmt.Printf("%v ", n-2)
				// } else {
				// 	fmt.Printf("%v ", n-i)
			}

		}
		fmt.Println()
	}
}
