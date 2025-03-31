package main

import "fmt"

/*
N = 5

A B C D E
A B C D
A B C
A B
A
*/
func main() {
	n := 5

	for i := 0; i < n; i++ {
		lt := 'A'
		for j := n; j > i; j-- {
			fmt.Printf("%v ", string(lt))
			lt++
		}
		fmt.Println()

	}
}
