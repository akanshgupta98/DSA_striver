package main

import (
	"fmt"
	"math"
)

/*
N =
*/

func checkPrime(n int) bool {

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
func main() {
	n := 1483
	fmt.Printf("The no %d is Prime: %v\n", n, checkPrime(n))
}
