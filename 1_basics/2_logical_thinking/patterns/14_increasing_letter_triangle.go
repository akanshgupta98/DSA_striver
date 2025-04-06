package main

import "fmt"

/*
N = 5
A
A B
A B C
A B C D
A B C D E
*/
func main() {
	n := 6

	for i := 0; i < n; i++ {
		lt := 'A'
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", string(lt))
			lt++
		}
		fmt.Println()

	}
}
