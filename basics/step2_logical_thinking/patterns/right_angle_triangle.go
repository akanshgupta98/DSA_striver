package main

import "fmt"

/*
N = 5
*
* *
* * *
* * * *
* * * * *
*/
func main() {
	N := 6
	for row := 0; row < N; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}

}
