package main

import "fmt"

/*
N=5

	   ********* i= 0, sp:0,9,0
		******* i = 1  1,7,1 -> (N*2-1 - 7)/2,7,(N*2-1 - 7)/2
		 *****  i=2,   2,5,2
		  ***   i=
		   *
*/
func main() {
	N := 10
	starCnt := (N * 2) - 1
	for i := 0; i < N; i++ {
		for k := 0; k < (((N*2)-1)-starCnt)/2; k++ {
			fmt.Printf(" ")
		}
		for k := 0; k < starCnt; k++ {
			fmt.Printf("*")
		}
		for k := 0; k < (((N*2)-1)-starCnt)/2; k++ {
			fmt.Printf(" ")
		}
		starCnt -= 2
		fmt.Println()

	}

}
