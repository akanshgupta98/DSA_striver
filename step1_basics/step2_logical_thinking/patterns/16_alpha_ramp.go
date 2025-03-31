package main

import "fmt"

/*
N = 5

A
BB
CCC
DDDD
EEEEE
*/
func main() {
	n := 5
	lt := 'A'
	for i := 0; i < n; i++ {

		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", string(lt))
		}
		lt++
		fmt.Println()

	}
}
