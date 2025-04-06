package main

import "fmt"

/*
N=5
*****
*   *
*   *
*   *
*****
*/
func main() {
	n := 6
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == n-1 || j == 0 || j == n-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Println()
	}

}
