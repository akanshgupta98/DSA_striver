package main

import "fmt"

func gcd(a, b int) int {

	if a == 0 {
		return b
	} else if b == 0 {
		return a
	}
	if a > b {
		return gcd(a-b, b)
	} else {
		return gcd(b-a, a)
	}

}
func main() {
	n1 := 12
	n2 := 18
	fmt.Printf("GCD of %d and %d is: %d\n", n1, n2, gcd(n1, n2))

}
