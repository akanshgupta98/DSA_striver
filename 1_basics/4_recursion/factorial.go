package main

import "fmt"

func calcFact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * calcFact(n-1)
}
func main() {
	n := 5
	fmt.Printf("Factorial of %d is: %d\n", n, calcFact(3))
}
