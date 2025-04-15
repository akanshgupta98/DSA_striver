package main

import "fmt"

func NSum(n int) int {
	if n == 1 {
		return 1
	}
	return n + NSum(n-1)
}
func main() {
	n := 6

	fmt.Printf("Sum of %d natural no is: %d\n", n, NSum(n))
}
