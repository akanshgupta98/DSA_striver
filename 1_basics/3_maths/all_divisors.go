package main

import (
	"fmt"
	"math"
)

/*
N = 36.
36 = [1,2,3,4,6,]
*/
func findDivisors(n int) (div []int) {
	// tmpMap := make(map[int]struct{})
	div = append(div, 1)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			// if _, ok := tmpMap[i]; !ok {
			// tmpMap[i] = struct{}{}
			div = append(div, i)
			if i != n/i {
				// tmpMap[n/i] = struct{}{}
				div = append(div, n/i)
			}

			// }

		}
	}
	div = append(div, n)
	return

}
func main() {
	n := 36
	div := findDivisors(n)
	fmt.Printf("Divisors of %d are: %v\n", n, div)

}
